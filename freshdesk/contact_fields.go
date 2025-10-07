package freshdesk

import (
	"context"
)

// List All Contact Fields
func (c *Client) ListContactFields(ctx context.Context) ([]*ContactField, error) {
	url := c.Endpoint("/admin/contact_fields")
	fields := []*ContactField{}
	err := c.DoGet(ctx, url, &fields)
	return fields, err
}

func (c *Client) CreateContactField(ctx context.Context, cf *ContactFieldCreate) (*ContactField, error) {
	url := c.Endpoint("/admin/contact_fields")
	result := &ContactField{}
	if err := c.DoPost(ctx, url, cf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetContactField View a Contact Field
// include: conversations, requester, company, stats
func (c *Client) GetContactField(ctx context.Context, fid int64) (*ContactField, error) {
	url := c.Endpoint("/admin/contact_fields/%d", fid)
	result := &ContactField{}
	err := c.DoGet(ctx, url, result)
	return result, err
}

func (c *Client) UpdateContactField(ctx context.Context, fid int64, field *ContactFieldUpdate) (*ContactField, error) {
	url := c.Endpoint("/admin/contact_fields/%d", fid)
	result := &ContactField{}
	if err := c.DoPut(ctx, url, field, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteContactField(ctx context.Context, fid int64) error {
	url := c.Endpoint("/admin/contact_fields/%d", fid)
	return c.DoDelete(ctx, url)
}
