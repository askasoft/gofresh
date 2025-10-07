package freshdesk

import (
	"context"
)

// List All Company Fields
func (c *Client) ListCompanyFields(ctx context.Context) ([]*CompanyField, error) {
	url := c.Endpoint("/admin/company_fields")
	fields := []*CompanyField{}
	err := c.DoGet(ctx, url, &fields)
	return fields, err
}

func (c *Client) CreateCompanyField(ctx context.Context, cf *CompanyFieldCreate) (*CompanyField, error) {
	url := c.Endpoint("/admin/company_fields")
	result := &CompanyField{}
	if err := c.DoPost(ctx, url, cf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCompanyField View a Company Field
// include: conversations, requester, company, stats
func (c *Client) GetCompanyField(ctx context.Context, fid int64) (*CompanyField, error) {
	url := c.Endpoint("/admin/company_fields/%d", fid)
	result := &CompanyField{}
	err := c.DoGet(ctx, url, result)
	return result, err
}

func (c *Client) UpdateCompanyField(ctx context.Context, fid int64, field *CompanyFieldUpdate) (*CompanyField, error) {
	url := c.Endpoint("/admin/company_fields/%d", fid)
	result := &CompanyField{}
	if err := c.DoPut(ctx, url, field, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteCompanyField(ctx context.Context, fid int64) error {
	url := c.Endpoint("/admin/company_fields/%d", fid)
	return c.DoDelete(ctx, url)
}
