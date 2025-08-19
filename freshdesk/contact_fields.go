package freshdesk

import (
	"context"
)

// List All Contact Fields
func (fd *Freshdesk) ListContactFields(ctx context.Context) ([]*ContactField, error) {
	url := fd.Endpoint("/admin/contact_fields")
	fields := []*ContactField{}
	err := fd.DoGet(ctx, url, &fields)
	return fields, err
}

func (fd *Freshdesk) CreateContactField(ctx context.Context, cf *ContactFieldCreate) (*ContactField, error) {
	url := fd.Endpoint("/admin/contact_fields")
	result := &ContactField{}
	if err := fd.DoPost(ctx, url, cf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetContactField View a Contact Field
// include: conversations, requester, company, stats
func (fd *Freshdesk) GetContactField(ctx context.Context, fid int64) (*ContactField, error) {
	url := fd.Endpoint("/admin/contact_fields/%d", fid)
	result := &ContactField{}
	err := fd.DoGet(ctx, url, result)
	return result, err
}

func (fd *Freshdesk) UpdateContactField(ctx context.Context, fid int64, field *ContactFieldUpdate) (*ContactField, error) {
	url := fd.Endpoint("/admin/contact_fields/%d", fid)
	result := &ContactField{}
	if err := fd.DoPut(ctx, url, field, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fd *Freshdesk) DeleteContactField(ctx context.Context, fid int64) error {
	url := fd.Endpoint("/admin/contact_fields/%d", fid)
	return fd.DoDelete(ctx, url)
}
