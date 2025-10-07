package fresh

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/askasoft/pango/doc/jsonx"
	"github.com/askasoft/pango/fsu"
	"github.com/askasoft/pango/iox"
	"github.com/askasoft/pango/log"
	"github.com/askasoft/pango/log/httplog"
	"github.com/askasoft/pango/net/httpx"
	"github.com/askasoft/pango/num"
	"github.com/askasoft/pango/ret"
)

const (
	contentTypeJSON = `application/json; charset="utf-8"`
)

type CustomRequest interface {
	RequestBody() (io.Reader, string, error)
}

type Client struct {
	Domain   string
	Apikey   string
	Username string
	Password string

	Transport http.RoundTripper
	Timeout   time.Duration
	Logger    log.Logger

	MaxRetries  int
	RetryAfter  time.Duration
	ShouldRetry func(error) bool // default retry on not canceled error or (status = 429 || (status >= 500 && status <= 599))
}

// Endpoint formats endpoint url
func (c *Client) Endpoint(format string, a ...any) string {
	return "https://" + c.Domain + "/api/v2" + fmt.Sprintf(format, a...)
}

func (c *Client) RetryForError(ctx context.Context, api func() error) (err error) {
	return ret.RetryForError(ctx, api, c.MaxRetries, c.Logger)
}

func (c *Client) authenticate(req *http.Request) {
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentTypeJSON)
	}

	if c.Apikey != "" {
		req.SetBasicAuth(c.Apikey, "X")
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}
}

func (c *Client) shouldRetry(err error) bool {
	sr := c.ShouldRetry
	if sr == nil {
		sr = shouldRetry
	}
	return sr(err)
}

func (c *Client) call(req *http.Request) (res *http.Response, err error) {
	client := &http.Client{
		Transport: c.Transport,
		Timeout:   c.Timeout,
	}

	res, err = httplog.TraceClientDo(c.Logger, client, req)
	if err != nil {
		if c.shouldRetry(err) {
			err = ret.NewRetryError(err, c.RetryAfter)
		}
	}

	return res, err
}

func (c *Client) authAndCall(req *http.Request) (*http.Response, error) {
	c.authenticate(req)
	return c.call(req)
}

func (c *Client) DoCall(req *http.Request, result any) error {
	_, err := c.doCall(req, result)
	return err
}

func (c *Client) doCall(req *http.Request, result any) (*http.Response, error) {
	res, err := c.authAndCall(req)
	if err != nil {
		return res, err
	}

	defer iox.DrainAndClose(res.Body)

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusNoContent {
		if result != nil {
			return res, decoder.Decode(result)
		}
		return res, nil
	}

	re := newResultError(res)
	if res.StatusCode != http.StatusNotFound {
		_ = decoder.Decode(re)
	}

	if c.shouldRetry(re) {
		s := res.Header.Get("Retry-After")
		n := num.Atoi(s)
		if n > 0 {
			re.RetryAfter = time.Second * time.Duration(n)
		} else {
			re.RetryAfter = c.RetryAfter
		}
	}

	return res, re
}

func (c *Client) DoGet(ctx context.Context, url string, result any) error {
	return c.RetryForError(ctx, func() error {
		return c.doGet(ctx, url, result)
	})
}

func (c *Client) doGet(ctx context.Context, url string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return c.DoCall(req, result)
}

func (c *Client) DoList(ctx context.Context, url string, lo ListOption, ap any) (next bool, err error) {
	err = c.RetryForError(ctx, func() error {
		next, err = c.doList(ctx, url, lo, ap)
		return err
	})
	return
}

func (c *Client) doList(ctx context.Context, url string, lo ListOption, result any) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	if lo != nil && !lo.IsNil() {
		q := lo.Values()
		req.URL.RawQuery = q.Encode()
	}

	res, err := c.doCall(req, result)
	if err != nil {
		return false, err
	}

	next := res.Header.Get("Link") != ""
	return next, nil
}

func (c *Client) DoPost(ctx context.Context, url string, source, result any) error {
	return c.RetryForError(ctx, func() error {
		return c.doPost(ctx, url, source, result)
	})
}

func (c *Client) doPost(ctx context.Context, url string, source, result any) error {
	buf, ct, err := buildRequest(source)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	return c.DoCall(req, result)
}

func (c *Client) DoPut(ctx context.Context, url string, source, result any) error {
	return c.RetryForError(ctx, func() error {
		return c.doPut(ctx, url, source, result)
	})
}

func (c *Client) doPut(ctx context.Context, url string, source, result any) error {
	buf, ct, err := buildRequest(source)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, buf)
	if err != nil {
		return err
	}
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	return c.DoCall(req, result)
}

func (c *Client) DoDelete(ctx context.Context, url string) error {
	return c.RetryForError(ctx, func() error {
		return c.doDelete(ctx, url)
	})
}

func (c *Client) doDelete(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	return c.DoCall(req, nil)
}

func (c *Client) DoDownload(ctx context.Context, url string) (buf []byte, err error) {
	err = c.RetryForError(ctx, func() error {
		buf, err = c.doDownload(ctx, url)
		return err
	})
	return
}

func (c *Client) doDownload(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.authAndCall(req)
	if err != nil {
		return nil, err
	}

	return copyResponse(res)
}

func (c *Client) DoSaveFile(ctx context.Context, url string, path string) error {
	return c.RetryForError(ctx, func() error {
		return c.doSaveFile(ctx, url, path)
	})
}

func (c *Client) doSaveFile(ctx context.Context, url string, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := c.authAndCall(req)
	if err != nil {
		return err
	}

	return saveResponse(res, path)
}

func (c *Client) DoDownloadNoAuth(ctx context.Context, url string) (buf []byte, err error) {
	err = c.RetryForError(ctx, func() error {
		buf, err = c.doDownloadNoAuth(ctx, url)
		return err
	})
	return
}

func (c *Client) doDownloadNoAuth(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.call(req)
	if err != nil {
		return nil, err
	}

	return copyResponse(res)
}

func (c *Client) DoSaveFileNoAuth(ctx context.Context, url string, path string) error {
	return c.RetryForError(ctx, func() error {
		return c.doSaveFileNoAuth(ctx, url, path)
	})
}

func (c *Client) doSaveFileNoAuth(ctx context.Context, url string, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := c.call(req)
	if err != nil {
		return err
	}

	return saveResponse(res, path)
}

func toString(o any) string {
	return jsonx.Prettify(o)
}

func addMultipartValues(mw *httpx.MultipartWriter, vs Values) error {
	return mw.WriteFields(url.Values(vs))
}

func addMultipartFiles(mw *httpx.MultipartWriter, fs Files) (err error) {
	for _, f := range fs {
		if f.Data() == nil {
			err = mw.WriteFile(f.Field(), f.File())
		} else {
			err = mw.WriteFileData(f.Field(), f.File(), f.Data())
		}
		if err != nil {
			return
		}
	}
	return
}

// buildRequest build a request, returns buffer, contentType, error
func buildRequest(a any) (io.Reader, string, error) {
	if a == nil {
		return nil, "", nil
	}

	if bb, ok := a.(CustomRequest); ok {
		return bb.RequestBody()
	}

	if wf, ok := a.(WithFiles); ok {
		fs := wf.Files()
		if len(fs) > 0 {
			vs := wf.Values()
			return BuildMultipartRequest(vs, fs)
		}
	}

	return BuildJSONRequest(a)
}

func BuildMultipartRequest(vs Values, fs Files) (io.Reader, string, error) {
	buf := &bytes.Buffer{}
	mw := httpx.NewMultipartWriter(buf)

	contentType := mw.FormDataContentType()

	if err := addMultipartValues(mw, vs); err != nil {
		return nil, "", err
	}
	if err := addMultipartFiles(mw, fs); err != nil {
		return nil, "", err
	}
	if err := mw.Close(); err != nil {
		return nil, "", err
	}

	return buf, contentType, nil
}

func BuildJSONRequest(a any) (io.Reader, string, error) {
	body, err := json.Marshal(a)
	if err != nil {
		return nil, "", err
	}

	buf := bytes.NewReader(body)
	return buf, contentTypeJSON, nil
}

func copyResponse(res *http.Response) ([]byte, error) {
	defer iox.DrainAndClose(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, newResultError(res)
	}

	buf := &bytes.Buffer{}
	_, err := iox.Copy(buf, res.Body)
	return buf.Bytes(), err
}

func saveResponse(res *http.Response, path string) error {
	defer iox.DrainAndClose(res.Body)

	if res.StatusCode != http.StatusOK {
		return newResultError(res)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0770); err != nil {
		return err
	}

	return fsu.WriteReader(path, res.Body, 0660)
}
