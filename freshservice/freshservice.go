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

func ParseDate(s string) (*Date, error) {
	return fresh.ParseDate(s)
}

func ParseTime(s string) (*Time, error) {
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

type Freshservice fresh.Fresh

func (fs *Freshservice) Endpoint(format string, a ...any) string {
	return (*fresh.Fresh)(fs).Endpoint(format, a...)
}

func (fs *Freshservice) DoGet(ctx context.Context, url string, result any) error {
	return (*fresh.Fresh)(fs).DoGet(ctx, url, result)
}

func (fs *Freshservice) DoList(ctx context.Context, url string, lo ListOption, result any) (bool, error) {
	return (*fresh.Fresh)(fs).DoList(ctx, url, lo, result)
}

func (fs *Freshservice) DoPost(ctx context.Context, url string, source, result any) error {
	return (*fresh.Fresh)(fs).DoPost(ctx, url, source, result)
}

func (fs *Freshservice) DoPut(ctx context.Context, url string, source, result any) error {
	return (*fresh.Fresh)(fs).DoPut(ctx, url, source, result)
}

func (fs *Freshservice) DoDelete(ctx context.Context, url string) error {
	return (*fresh.Fresh)(fs).DoDelete(ctx, url)
}

func (fs *Freshservice) Download(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Fresh)(fs).DoDownload(ctx, url)
}

func (fs *Freshservice) SaveFile(ctx context.Context, url string, path string) error {
	return (*fresh.Fresh)(fs).DoSaveFile(ctx, url, path)
}

func (fs *Freshservice) DownloadNoAuth(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Fresh)(fs).DoDownloadNoAuth(ctx, url)
}

func (fs *Freshservice) SaveFileNoAuth(ctx context.Context, url string, path string) error {
	return (*fresh.Fresh)(fs).DoSaveFileNoAuth(ctx, url, path)
}

func (fs *Freshservice) DownloadAttachment(ctx context.Context, aid int64) ([]byte, error) {
	url := fs.Endpoint("/attachments/%d", aid)
	return fs.Download(ctx, url)
}

func (fs *Freshservice) SaveAttachment(ctx context.Context, aid int64, path string) error {
	url := fs.Endpoint("/attachments/%d", aid)
	return fs.SaveFile(ctx, url, path)
}

// GetAgentTicketURL return a permlink for agent ticket URL
func (fs *Freshservice) GetAgentTicketURL(tid int64) string {
	return GetAgentTicketURL(fs.Domain, tid)
}

// GetSolutionArticleURL return a permlink for solution article URL
func (fs *Freshservice) GetSolutionArticleURL(aid int64) string {
	return GetSolutionArticleURL(fs.Domain, aid)
}

// GetHelpdeskAttachmentURL return a permlink for helpdesk attachment/avator URL
func (fs *Freshservice) GetHelpdeskAttachmentURL(aid int64) string {
	return GetHelpdeskAttachmentURL(fs.Domain, aid)
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
