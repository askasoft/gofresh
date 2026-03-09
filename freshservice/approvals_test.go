package freshservice

import (
	"testing"
)

func TestIterApprovals(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	itcnt := 0
	err := fs.IterApprovals(ctxbg, &ListApprovalsOption{
		Parent: "ticket",
		Status: ApprovalStatusRequested.String(),
	}, func(a *Approval) error {
		itcnt++
		tlog.Debugf("Iterate Approval #%d: %s", itcnt, a)
		return nil
	})
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	tlog.Infof("Iterate %d approvals", itcnt)
}
