package freshservice

import (
	"testing"
)

func TestWorkspaces(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	itcnt := 0
	err := fs.IterWorkspaces(ctxbg, nil, func(w *Workspace) error {
		itcnt++
		tlog.Debugf("Iterate workspace #%d: %s", w.ID, w.Name)
		return nil
	})
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	tlog.Infof("Iterate %d workspaces", itcnt)
}
