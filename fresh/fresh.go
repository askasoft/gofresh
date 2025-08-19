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

type Fresh struct {
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
func (fresh *Fresh) Endpoint(format string, a ...any) string {
	return "https://" + fresh.Domain + "/api/v2" + fmt.Sprintf(format, a...)
}

func (fresh *Fresh) RetryForError(ctx context.Context, api func() error) (err error) {
	return ret.RetryForError(ctx, api, fresh.MaxRetries, fresh.Logger)
}

func (fresh *Fresh) authenticate(req *http.Request) {
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentTypeJSON)
	}

	if fresh.Apikey != "" {
		req.SetBasicAuth(fresh.Apikey, "X")
	} else {
		req.SetBasicAuth(fresh.Username, fresh.Password)
	}
}

func (fresh *Fresh) call(req *http.Request) (res *http.Response, err error) {
	client := &http.Client{
		Transport: fresh.Transport,
		Timeout:   fresh.Timeout,
	}

	if log := fresh.Logger; log != nil {
		log.Debugf("%s %s", req.Method, req.URL)
	}

	rid := httplog.TraceHttpRequest(fresh.Logger, req)

	res, err = client.Do(req)
	if err != nil {
		fsr := fresh.ShouldRetry
		if fsr == nil {
			fsr = shouldRetry
		}
		if fsr(err) {
			err = ret.NewRetryError(err, fresh.RetryAfter)
		}
		return res, err
	}

	httplog.TraceHttpResponse(fresh.Logger, res, rid)
	return res, nil
}

func (fresh *Fresh) authAndCall(req *http.Request) (res *http.Response, err error) {
	fresh.authenticate(req)
	return fresh.call(req)
}

func (fresh *Fresh) doCall(req *http.Request, result any) error {
	res, err := fresh.authAndCall(req)
	if err != nil {
		return err
	}

	return fresh.decodeResponse(res, result)
}

func (fresh *Fresh) decodeResponse(res *http.Response, result any) error {
	defer iox.DrainAndClose(res.Body)

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusNoContent {
		if result != nil {
			return decoder.Decode(result)
		}
		return nil
	}

	re := newResultError(res)
	if res.StatusCode != http.StatusNotFound {
		_ = decoder.Decode(re)
	}

	fsr := fresh.ShouldRetry
	if fsr == nil {
		fsr = shouldRetry
	}

	if fsr(re) {
		s := res.Header.Get("Retry-After")
		n := num.Atoi(s)
		if n > 0 {
			re.RetryAfter = time.Second * time.Duration(n)
		} else {
			re.RetryAfter = fresh.RetryAfter
		}
	}

	return re
}

func (fresh *Fresh) DoGet(ctx context.Context, url string, result any) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doGet(ctx, url, result)
	})
}

func (fresh *Fresh) doGet(ctx context.Context, url string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return fresh.doCall(req, result)
}

func (fresh *Fresh) DoList(ctx context.Context, url string, lo ListOption, ap any) (next bool, err error) {
	err = fresh.RetryForError(ctx, func() error {
		next, err = fresh.doList(ctx, url, lo, ap)
		return err
	})
	return
}

func (fresh *Fresh) doList(ctx context.Context, url string, lo ListOption, result any) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	if lo != nil && !lo.IsNil() {
		q := lo.Values()
		req.URL.RawQuery = q.Encode()
	}

	res, err := fresh.authAndCall(req)
	if err != nil {
		return false, err
	}

	err = fresh.decodeResponse(res, result)
	if err != nil {
		return false, err
	}

	next := res.Header.Get("Link") != ""
	return next, nil
}

func (fresh *Fresh) DoPost(ctx context.Context, url string, source, result any) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doPost(ctx, url, source, result)
	})
}

func (fresh *Fresh) doPost(ctx context.Context, url string, source, result any) error {
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

	return fresh.doCall(req, result)
}

func (fresh *Fresh) DoPut(ctx context.Context, url string, source, result any) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doPut(ctx, url, source, result)
	})
}

func (fresh *Fresh) doPut(ctx context.Context, url string, source, result any) error {
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

	return fresh.doCall(req, result)
}

func (fresh *Fresh) DoDelete(ctx context.Context, url string) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doDelete(ctx, url)
	})
}

func (fresh *Fresh) doDelete(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	return fresh.doCall(req, nil)
}

func (fresh *Fresh) DoDownload(ctx context.Context, url string) (buf []byte, err error) {
	err = fresh.RetryForError(ctx, func() error {
		buf, err = fresh.doDownload(ctx, url)
		return err
	})
	return
}

func (fresh *Fresh) doDownload(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := fresh.authAndCall(req)
	if err != nil {
		return nil, err
	}

	return copyResponse(res)
}

func (fresh *Fresh) DoSaveFile(ctx context.Context, url string, path string) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doSaveFile(ctx, url, path)
	})
}

func (fresh *Fresh) doSaveFile(ctx context.Context, url string, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := fresh.authAndCall(req)
	if err != nil {
		return err
	}

	return saveResponse(res, path)
}

func (fresh *Fresh) DoDownloadNoAuth(ctx context.Context, url string) (buf []byte, err error) {
	err = fresh.RetryForError(ctx, func() error {
		buf, err = fresh.doDownloadNoAuth(ctx, url)
		return err
	})
	return
}

func (fresh *Fresh) doDownloadNoAuth(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := fresh.call(req)
	if err != nil {
		return nil, err
	}

	return copyResponse(res)
}

func (fresh *Fresh) DoSaveFileNoAuth(ctx context.Context, url string, path string) error {
	return fresh.RetryForError(ctx, func() error {
		return fresh.doSaveFileNoAuth(ctx, url, path)
	})
}

func (fresh *Fresh) doSaveFileNoAuth(ctx context.Context, url string, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := fresh.call(req)
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
	if err := os.MkdirAll(dir, os.FileMode(0770)); err != nil {
		return err
	}

	return fsu.WriteReader(path, res.Body, fsu.FileMode(0660))
}
