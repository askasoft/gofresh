package freshdesk

import (
	"testing"
)

func TestCompanyFieldsAPIs(t *testing.T) {
	fd := testNewFreshdesk(t)
	if fd == nil {
		return
	}

	tfc := &CompanyFieldCreate{
		Label: "testfieldlabel",
		Type:  CustomFieldTypeCustomText,
	}

	ctf, err := fd.CreateCompanyField(ctxbg, tfc)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	fd.Logger.Debug(ctf)

	tfu := &CompanyFieldUpdate{
		Label: "testfieldlabelforupd",
	}

	utf, err := fd.UpdateCompanyField(ctxbg, ctf.ID, tfu)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fd.Logger.Debug(utf)
	}

	gtfr, err := fd.GetCompanyField(ctxbg, ctf.ID)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fd.Logger.Debug(gtfr)
	}

	err = fd.DeleteCompanyField(ctxbg, ctf.ID)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
}

func TestListCompanyFieldsAPIs(t *testing.T) {
	fd := testNewFreshdesk(t)
	if fd == nil {
		return
	}

	tfs, err := fd.ListCompanyFields(ctxbg)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	fd.Logger.Debug(tfs)
}
