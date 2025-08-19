package freshdesk

import (
	"context"
)

// List All Company Fields
func (fd *Freshdesk) ListCompanyFields(ctx context.Context) ([]*CompanyField, error) {
	url := fd.Endpoint("/admin/company_fields")
	fields := []*CompanyField{}
	err := fd.DoGet(ctx, url, &fields)
	return fields, err
}

func (fd *Freshdesk) CreateCompanyField(ctx context.Context, cf *CompanyFieldCreate) (*CompanyField, error) {
	url := fd.Endpoint("/admin/company_fields")
	result := &CompanyField{}
	if err := fd.DoPost(ctx, url, cf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCompanyField View a Company Field
// include: conversations, requester, company, stats
func (fd *Freshdesk) GetCompanyField(ctx context.Context, fid int64) (*CompanyField, error) {
	url := fd.Endpoint("/admin/company_fields/%d", fid)
	result := &CompanyField{}
	err := fd.DoGet(ctx, url, result)
	return result, err
}

func (fd *Freshdesk) UpdateCompanyField(ctx context.Context, fid int64, field *CompanyFieldUpdate) (*CompanyField, error) {
	url := fd.Endpoint("/admin/company_fields/%d", fid)
	result := &CompanyField{}
	if err := fd.DoPut(ctx, url, field, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (fd *Freshdesk) DeleteCompanyField(ctx context.Context, fid int64) error {
	url := fd.Endpoint("/admin/company_fields/%d", fid)
	return fd.DoDelete(ctx, url)
}
