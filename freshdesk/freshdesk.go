package freshdesk

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

	DefaultFieldTypeName        = "default_name"
	DefaultFieldTypeNote        = "default_note"
	DefaultFieldTypeSubject     = "default_subject"
	DefaultFieldTypeAgent       = "default_agent"
	DefaultFieldTypeRequester   = "default_requester"
	DefaultFieldTypeGroup       = "default_group"
	DefaultFieldTypeCompany     = "default_company"
	DefaultFieldTypeProduct     = "default_product"
	DefaultFieldTypePriority    = "default_priority"
	DefaultFieldTypeSource      = "default_source"
	DefaultFieldTypeStatus      = "default_status"
	DefaultFieldTypeDescription = "default_description"
	DefaultFieldTypeTicketType  = "default_ticket_type"

	CustomFieldTypeCustomDate      = "custom_date"
	CustomFieldTypeCustomDateTime  = "custom_date_time"
	CustomFieldTypeCustomDropdown  = "custom_dropdown"
	CustomFieldTypeCustomParagraph = "custom_paragraph"
	CustomFieldTypeCustomText      = "custom_text"
	CustomFieldTypeCustomCheckbox  = "custom_checkbox"
	CustomFieldTypeCustomNumber    = "custom_number"
	CustomFieldTypeCustomDecimal   = "custom_decimal"
	CustomFieldTypeCustomFile      = "custom_file"
	CustomFieldTypeNestedField     = "nested_field"
)

type FilterOption struct {
	Query string
	Page  int
}

func (fo *FilterOption) IsNil() bool {
	return fo == nil
}

func (fo *FilterOption) Values() Values {
	q := Values{}
	q.SetString("query", fo.Query)
	q.SetInt("page", fo.Page)
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

type Freshdesk fresh.Fresh

func (fd *Freshdesk) Endpoint(format string, a ...any) string {
	return (*fresh.Fresh)(fd).Endpoint(format, a...)
}

func (fd *Freshdesk) DoGet(ctx context.Context, url string, result any) error {
	return (*fresh.Fresh)(fd).DoGet(ctx, url, result)
}

func (fd *Freshdesk) DoList(ctx context.Context, url string, lo ListOption, result any) (bool, error) {
	return (*fresh.Fresh)(fd).DoList(ctx, url, lo, result)
}

func (fd *Freshdesk) DoPost(ctx context.Context, url string, source, result any) error {
	return (*fresh.Fresh)(fd).DoPost(ctx, url, source, result)
}

func (fd *Freshdesk) DoPut(ctx context.Context, url string, source, result any) error {
	return (*fresh.Fresh)(fd).DoPut(ctx, url, source, result)
}

func (fd *Freshdesk) DoDelete(ctx context.Context, url string) error {
	return (*fresh.Fresh)(fd).DoDelete(ctx, url)
}

func (fd *Freshdesk) Download(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Fresh)(fd).DoDownload(ctx, url)
}

func (fd *Freshdesk) SaveFile(ctx context.Context, url string, path string) error {
	return (*fresh.Fresh)(fd).DoSaveFile(ctx, url, path)
}

func (fd *Freshdesk) DownloadNoAuth(ctx context.Context, url string) ([]byte, error) {
	return (*fresh.Fresh)(fd).DoDownloadNoAuth(ctx, url)
}

func (fd *Freshdesk) SaveFileNoAuth(ctx context.Context, url string, path string) error {
	return (*fresh.Fresh)(fd).DoSaveFileNoAuth(ctx, url, path)
}

func (fd *Freshdesk) DeleteAttachment(ctx context.Context, aid int64) error {
	url := fd.Endpoint("/attachments/%d", aid)
	return fd.DoDelete(ctx, url)
}

// GetJob get job detail
func (fd *Freshdesk) GetJob(ctx context.Context, jid string) (*Job, error) {
	url := fd.Endpoint("/jobs/%s", jid)
	job := &Job{}
	err := fd.DoGet(ctx, url, job)
	return job, err
}

// GetAgentTicketURL return a permlink for agent ticket URL
func (fd *Freshdesk) GetAgentTicketURL(tid int64) string {
	return GetAgentTicketURL(fd.Domain, tid)
}

// GetSolutionArticleURL return a permlink for solution article URL
func (fd *Freshdesk) GetSolutionArticleURL(aid int64) string {
	return GetSolutionArticleURL(fd.Domain, aid)
}

// GetHelpdeskAttachmentURL return a permlink for helpdesk attachment/avator URL
func (fd *Freshdesk) GetHelpdeskAttachmentURL(aid int64) string {
	return GetHelpdeskAttachmentURL(fd.Domain, aid)
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
