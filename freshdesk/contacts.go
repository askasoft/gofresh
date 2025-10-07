package freshdesk

import (
	"context"
	"net/url"
)

// ---------------------------------------------------
// Contact

type ContactState string

const (
	ContactStateBlocked    ContactState = "blocked"
	ContactStateDeleted    ContactState = "deleted"
	ContactStateUnverified ContactState = "unverified"
	ContactStateVerified   ContactState = "verified"
)

type ListContactsOption struct {
	Email            string
	Mobile           string
	Phone            string
	UniqueExternalID string
	CompanyID        int64
	UpdatedSince     Time
	State            ContactState // [blocked/deleted/unverified/verified]
	Page             int
	PerPage          int
}

func (lco *ListContactsOption) IsNil() bool {
	return lco == nil
}

func (lco *ListContactsOption) Values() Values {
	q := Values{}
	q.SetString("email", lco.Email)
	q.SetString("mobile", lco.Mobile)
	q.SetString("phone", lco.Phone)
	q.SetString("unique_external_id", lco.UniqueExternalID)
	q.SetInt64("company_id", lco.CompanyID)
	q.SetString("state", (string)(lco.State))
	q.SetTime("updated_since", lco.UpdatedSince)
	q.SetInt("page", lco.Page)
	q.SetInt("per_page", lco.PerPage)
	return q
}

type FilterContactsOption = FilterOption

type FilterContactsResult struct {
	Total   int        `json:"total"`
	Results []*Contact `json:"results"`
}

func (c *Client) CreateContact(ctx context.Context, contact *ContactCreate) (*Contact, error) {
	url := c.Endpoint("/contacts")
	result := &Contact{}
	if err := c.DoPost(ctx, url, contact, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateContact(ctx context.Context, cid int64, contact *ContactUpdate) (*Contact, error) {
	url := c.Endpoint("/contacts/%d", cid)
	result := &Contact{}
	if err := c.DoPut(ctx, url, contact, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetContact(ctx context.Context, cid int64) (*Contact, error) {
	url := c.Endpoint("/contacts/%d", cid)
	contact := &Contact{}
	err := c.DoGet(ctx, url, contact)
	return contact, err
}

func (c *Client) DeleteContact(ctx context.Context, cid int64) error {
	url := c.Endpoint("/contacts/%d", cid)
	return c.DoDelete(ctx, url)
}

func (c *Client) HardDeleteContact(ctx context.Context, cid int64, force ...bool) error {
	url := c.Endpoint("/contacts/%d/hard_delete", cid)
	if len(force) > 0 && force[0] {
		url += "?force=true"
	}
	return c.DoDelete(ctx, url)
}

func (c *Client) ListContacts(ctx context.Context, lco *ListContactsOption) ([]*Contact, bool, error) {
	url := c.Endpoint("/contacts")
	contacts := []*Contact{}
	next, err := c.DoList(ctx, url, lco, &contacts)
	return contacts, next, err
}

func (c *Client) IterContacts(ctx context.Context, lco *ListContactsOption, icf func(*Contact) error) error {
	if lco == nil {
		lco = &ListContactsOption{}
	}
	if lco.Page < 1 {
		lco.Page = 1
	}
	if lco.PerPage < 1 {
		lco.PerPage = 100
	}

	for {
		contacts, next, err := c.ListContacts(ctx, lco)
		if err != nil {
			return err
		}
		for _, c := range contacts {
			if err = icf(c); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lco.Page++
	}
	return nil
}

func (c *Client) SearchContacts(ctx context.Context, keyword string) ([]*User, error) {
	url := c.Endpoint("/contacts/autocomplete?term=%s", url.QueryEscape(keyword))
	contacts := []*User{}
	err := c.DoGet(ctx, url, &contacts)
	return contacts, err
}

// FilterContacts Use custom contact fields that you have created in your account to filter through the contacts and get a list of contacts matching the specified contact fields.
// Format: "(contact_field:integer OR contact_field:'string') AND contact_field:boolean"
// See: https://developers.freshdesk.com/api/#filter_contacts
func (c *Client) FilterContacts(ctx context.Context, fco *FilterContactsOption) ([]*Contact, int, error) {
	url := c.Endpoint("/search/contacts")
	fcr := &FilterContactsResult{}
	_, err := c.DoList(ctx, url, fco, fcr)
	return fcr.Results, fcr.Total, err
}

func (c *Client) RestoreContact(ctx context.Context, cid int64) error {
	url := c.Endpoint("/contacts/%d/restore", cid)
	return c.DoPut(ctx, url, nil, nil)
}

func (c *Client) InviteContact(ctx context.Context, cid int64) error {
	url := c.Endpoint("/contacts/%d/send_invite", cid)
	return c.DoPut(ctx, url, nil, nil)
}

func (c *Client) MergeContacts(ctx context.Context, cm *ContactsMerge) error {
	url := c.Endpoint("/contacts/merge")
	return c.DoPost(ctx, url, nil, nil)
}

// ExportContacts return a job id, call GetExportedContactsURL() to get the job detail
func (c *Client) ExportContacts(ctx context.Context, defaultFields, customFields []string) (string, error) {
	url := c.Endpoint("/contacts/export")
	opt := &ExportOption{
		Fields: &ExportFields{
			DefaultFields: defaultFields,
			CustomFields:  customFields,
		},
	}
	job := &Job{}
	err := c.DoPost(ctx, url, opt, job)
	return job.ID, err
}

// GetExportedContactsURL get the exported contacts url
func (c *Client) GetExportedContactsURL(ctx context.Context, jid string) (*Job, error) {
	url := c.Endpoint("/contacts/export/%s", jid)
	job := &Job{}
	err := c.DoGet(ctx, url, job)
	return job, err
}

func (c *Client) MakeAgent(ctx context.Context, cid int64, agent *Agent) (*Contact, error) {
	url := c.Endpoint("/contacts/%d/make_agent", cid)
	result := &Contact{}
	if err := c.DoPut(ctx, url, agent, result); err != nil {
		return nil, err
	}
	return result, nil
}
