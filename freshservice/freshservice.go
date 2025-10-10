package freshservice

import (
	"context"
	"fmt"
	"io"

	"github.com/askasoft/gofresh/fresh"
	"github.com/askasoft/pango/doc/jsonx"
)

type FieldError = fresh.FieldError
type ResultError = fresh.ResultError
type Date = fresh.Date
type Time = fresh.Time
type TimeSpent = fresh.TimeSpent
type Attachment = fresh.Attachment
type Attachments = fresh.Attachments
type ListOption = fresh.ListOption
type PageOption = fresh.PageOption
type File = fresh.File
type Files = fresh.Files
type WithFiles = fresh.WithFiles
type Values = fresh.Values

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

type FilterOption struct {
	Query   string
	Page    int
	PerPage int
}

func (fo *FilterOption) IsNil() bool {
	return fo == nil
}

func (fo *FilterOption) Values() Values {
	q := Values{}
	q.SetString("query", fo.Query)
	q.SetInt("page", fo.Page)
	q.SetInt("per_page", fo.PerPage)
	return q
}

func ParseDate(s string) (Date, error) {
	return fresh.ParseDate(s)
}

func ParseTime(s string) (Time, error) {
	return fresh.ParseTime(s)
}

func ParseTimeSpent(s string) (TimeSpent, error) {
	return fresh.ParseTimeSpent(s)
}

func NewAttachment(file string, data ...[]byte) *Attachment {
	return fresh.NewAttachment(file, data...)
}

func toString(o any) string {
	return jsonx.Prettify(o)
}

// alias for Client
type Freshservice = Client

type Client fresh.Client

func (c *Client) Endpoint(format string, a ...any) string {
	return (*fresh.Client)(c).Endpoint(format, a...)
}

func (c *Client) DoGet(ctx context.Context, url string, result any) error {
	return (*fresh.Client)(c).DoGet(ctx, url, result)
}

func (c *Client) DoList(ctx context.Context, url string, lo ListOption, result any) (bool, error) {
	return (*fresh.Client)(c).DoList(ctx, url, lo, result)
}

func (c *Client) DoPost(ctx context.Context, url string, source, result any) error {
	return (*fresh.Client)(c).DoPost(ctx, url, source, result)
}

func (c *Client) DoPut(ctx context.Context, url string, source, result any) error {
	return (*fresh.Client)(c).DoPut(ctx, url, source, result)
}

func (c *Client) DoDelete(ctx context.Context, url string) error {
	return (*fresh.Client)(c).DoDelete(ctx, url)
}

func (c *Client) DoCopyFile(ctx context.Context, url string, w io.Writer) error {
	return (*fresh.Client)(c).DoCopyFile(ctx, url, w)
}

func (c *Client) DoReadFile(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Client)(c).DoReadFile(ctx, url)
}

func (c *Client) DoSaveFile(ctx context.Context, url string, path string) error {
	return (*fresh.Client)(c).DoSaveFile(ctx, url, path)
}

func (c *Client) DoCopyFileNoAuth(ctx context.Context, url string, w io.Writer) error {
	return (*fresh.Client)(c).DoCopyFileNoAuth(ctx, url, w)
}

func (c *Client) DoReadFileNoAuth(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Client)(c).DoReadFileNoAuth(ctx, url)
}

func (c *Client) DoSaveFileNoAuth(ctx context.Context, url string, path string) error {
	return (*fresh.Client)(c).DoSaveFileNoAuth(ctx, url, path)
}

func (c *Client) CopyAttachment(ctx context.Context, aid int64, w io.Writer) error {
	url := c.Endpoint("/attachments/%d", aid)
	return c.DoCopyFile(ctx, url, w)
}

func (c *Client) ReadAttachment(ctx context.Context, aid int64) ([]byte, error) {
	url := c.Endpoint("/attachments/%d", aid)
	return c.DoReadFile(ctx, url)
}

func (c *Client) SaveAttachment(ctx context.Context, aid int64, path string) error {
	url := c.Endpoint("/attachments/%d", aid)
	return c.DoSaveFile(ctx, url, path)
}

// unsupported by Freshservice API
// func (c *Client) DeleteAttachment(ctx context.Context, aid int64) error {
// 	url := c.Endpoint("/attachments/%d", aid)
// 	return c.DoDelete(ctx, url)
// }

// GetAgentTicketURL return a permlink for agent ticket URL
func (c *Client) GetAgentTicketURL(tid int64) string {
	return GetAgentTicketURL(c.Domain, tid)
}

// GetSolutionArticleURL return a permlink for solution article URL
func (c *Client) GetSolutionArticleURL(aid int64) string {
	return GetSolutionArticleURL(c.Domain, aid)
}

// GetHelpdeskAttachmentURL return a permlink for helpdesk attachment/avator URL
func (c *Client) GetHelpdeskAttachmentURL(aid int64) string {
	return GetHelpdeskAttachmentURL(c.Domain, aid)
}

// GetAgentTicketURL return a permlink for agent ticket URL
func GetAgentTicketURL(domain string, tid int64) string {
	return fmt.Sprintf("https://%s/a/tickets/%d", domain, tid)
}

// GetSolutionArticleURL return a permlink for solution article URL
func GetSolutionArticleURL(domain string, aid int64) string {
	return fmt.Sprintf("https://%s/support/solutions/articles/%d", domain, aid)
}

// GetHelpdeskAttachmentURL return a permlink for helpdesk attachment/avator URL
func GetHelpdeskAttachmentURL(domain string, aid int64) string {
	return fmt.Sprintf("https://%s/helpdesk/attachments/%d", domain, aid)
}
