package freshdesk

import (
	"testing"
)

func TestContactFieldsAPIs(t *testing.T) {
	fd := testNewFreshdesk(t)
	if fd == nil {
		return
	}

	tfc := &ContactFieldCreate{
		Label:                 "testfieldlabel",
		LabelForCustomers:     "testfieldlabelforcustomers",
		Type:                  CustomFieldTypeCustomText,
		CustomersCanEdit:      true,
		DisplayedForCustomers: true,
	}

	ctf, err := fd.CreateContactField(ctxbg, tfc)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	fd.Logger.Debug(ctf)

	tfu := &ContactFieldUpdate{
		LabelForCustomers: "testfieldlabelforcustomersupd",
	}

	utf, err := fd.UpdateContactField(ctxbg, ctf.ID, tfu)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fd.Logger.Debug(utf)
	}

	gtfr, err := fd.GetContactField(ctxbg, ctf.ID)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fd.Logger.Debug(gtfr)
	}

	err = fd.DeleteContactField(ctxbg, ctf.ID)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
}

func TestListContactFieldsAPIs(t *testing.T) {
	fd := testNewFreshdesk(t)
	if fd == nil {
		return
	}

	tfs, err := fd.ListContactFields(ctxbg)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	fd.Logger.Debug(tfs)
}
