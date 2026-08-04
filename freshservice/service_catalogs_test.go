package freshservice

import (
	"fmt"
	"testing"
	"time"
)

func TestIterServiceCategories(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	cnt := 0
	err := fs.IterServiceCategories(ctxbg, nil, func(sc *ServiceCategory) error {
		cnt++
		tlog.Infof("[%d] Service Category #%d: %s", cnt, sc.ID, sc.Name)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestIterServiceItems(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	cnt := 0
	err := fs.IterServiceItems(ctxbg, nil, func(si *ServiceItem) error {
		si, err := fs.GetServiceItem(ctxbg, si.DisplayID)
		if err != nil {
			return err
		}

		cnt++
		tlog.Infof("[%d] Service Item #%d: %s", cnt, si.DisplayID, si.Name)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestIterServiceCatalog(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	cnt := 0
	err := fs.IterServiceCategories(ctxbg, nil, func(sc *ServiceCategory) error {
		tlog.Infof("Service Category #%d: %s", sc.ID, sc.Name)

		return fs.IterServiceItems(ctxbg, &ListServiceItemsOption{CategoryID: sc.ID}, func(si *ServiceItem) error {
			si, err := fs.GetServiceItem(ctxbg, si.DisplayID)
			if err != nil {
				return err
			}

			cnt++
			tlog.Infof("[%d] Service Item #%d: %s (%v)", cnt, si.DisplayID, si.Name, si.Cost)
			return nil
		})
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetServiceItem(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	si, err := fs.GetServiceItem(ctxbg, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(si)
}

func TestServiceItems(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	scs, _, err := fs.ListServiceCategories(ctxbg, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(scs) == 0 {
		return
	}

	sfx := time.Now().Format("20060102150405")
	sic := &ServiceItemCreate{
		CategoryID:             scs[0].ID,
		Name:                   "Test Item " + sfx,
		ShortDescription:       "test short " + sfx,
		Description:            "test description " + sfx,
		Visibility:             ServiceItemDraft,
		AgentGroupVisibility:   AgentGroupVisibilityGroups,
		AgentGroupVisibilities: &AgentGroupVisibilities{GroupID: []int64{17000334151, 17000384862}},
	}
	res1, err := fs.CreateServiceItem(ctxbg, sic)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := fs.DeleteServiceItem(ctxbg, res1.DisplayID)
		if err != nil {
			t.Fatal(err)
		}
	}()

	_, err = fs.GetServiceItem(ctxbg, res1.DisplayID)
	if err != nil {
		t.Error(err)
	}

	sfx += "-2"
	siu := &ServiceItemUpdate{
		CategoryID:       scs[0].ID,
		Name:             "Test Item " + sfx,
		ShortDescription: "test short " + sfx,
		Description:      "test description " + sfx,
		Visibility:       ServiceItemDraft,
	}

	_, err = fs.UpdateServiceItem(ctxbg, res1.DisplayID, siu)
	if err != nil {
		t.Error(err)
	}
}
