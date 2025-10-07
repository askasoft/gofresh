package freshdesk

import (
	"context"
	"net/url"
)

// ---------------------------------------------------
// Company

type ListCompaniesOption = PageOption

func (c *Client) CreateCompany(ctx context.Context, company *CompanyCreate) (*Company, error) {
	url := c.Endpoint("/companies")
	result := &Company{}
	if err := c.DoPost(ctx, url, company, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetCompany(ctx context.Context, cid int64) (*Company, error) {
	url := c.Endpoint("/companies/%d", cid)
	result := &Company{}
	err := c.DoGet(ctx, url, result)
	return result, err
}

func (c *Client) ListCompanies(ctx context.Context, lco *ListCompaniesOption) ([]*Company, bool, error) {
	url := c.Endpoint("/companies")
	result := []*Company{}
	next, err := c.DoList(ctx, url, lco, &result)
	return result, next, err
}

func (c *Client) IterCompanies(ctx context.Context, lco *ListCompaniesOption, icf func(*Company) error) error {
	if lco == nil {
		lco = &ListCompaniesOption{}
	}
	if lco.Page < 1 {
		lco.Page = 1
	}
	if lco.PerPage < 1 {
		lco.PerPage = 100
	}

	for {
		companies, next, err := c.ListCompanies(ctx, lco)
		if err != nil {
			return err
		}
		for _, c := range companies {
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

// Search Companies
// Search for a company using its name.
// Note:
// 1. The search is case insensitive.
// 2. You cannot search with a substring. For example, a company "Acme Corporation" can be looked up using "acme", "Ac", "Corporation" and "Co". But it will not be returned when you search for "cme" or "orporation".
func (c *Client) SearchCompanies(ctx context.Context, name string) ([]*Company, error) {
	url := c.Endpoint("/companies/autocomplete?name=%s", url.QueryEscape(name))
	result := &companyResult{}
	err := c.DoGet(ctx, url, result)
	return result.Companies, err
}

func (c *Client) UpdateCompany(ctx context.Context, cid int64, company *CompanyUpdate) (*Company, error) {
	url := c.Endpoint("/companies/%d", cid)
	result := &Company{}
	if err := c.DoPut(ctx, url, company, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteCompany(ctx context.Context, cid int64) error {
	url := c.Endpoint("/companies/%d", cid)
	return c.DoDelete(ctx, url)
}

// ExportCompanies return a job id, call GetExportedCompaniesURL() to get the job detail
func (c *Client) ExportCompanies(ctx context.Context, defaultFields, customFields []string) (string, error) {
	url := c.Endpoint("/companies/export")
	opt := &ExportOption{
		Fields: &ExportFields{
			DefaultFields: defaultFields,
			CustomFields:  customFields,
		},
	}
	job := &Job{}
	err := c.DoPost(ctx, url, opt, &job)
	return job.ID, err
}

// GetExportedCompaniesURL get the exported companies url
func (c *Client) GetExportedCompaniesURL(ctx context.Context, jid string) (*Job, error) {
	url := c.Endpoint("/companies/export/%s", jid)
	job := &Job{}
	err := c.DoGet(ctx, url, job)
	return job, err
}
