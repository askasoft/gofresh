package freshservice

import (
	"context"
	"fmt"

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

func (c *Client) Download(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Client)(c).DoDownload(ctx, url)
}

func (c *Client) SaveFile(ctx context.Context, url string, path string) error {
	return (*fresh.Client)(c).DoSaveFile(ctx, url, path)
}

func (c *Client) DownloadNoAuth(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Client)(c).DoDownloadNoAuth(ctx, url)
}

func (c *Client) SaveFileNoAuth(ctx context.Context, url string, path string) error {
	return (*fresh.Client)(c).DoSaveFileNoAuth(ctx, url, path)
}

func (c *Client) DownloadAttachment(ctx context.Context, aid int64) ([]byte, error) {
	url := c.Endpoint("/attachments/%d", aid)
	return c.Download(ctx, url)
}

func (c *Client) SaveAttachment(ctx context.Context, aid int64, path string) error {
	url := c.Endpoint("/attachments/%d", aid)
	return c.SaveFile(ctx, url, path)
}

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
