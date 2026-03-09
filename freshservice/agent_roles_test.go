package freshservice

import (
	"testing"
)

func TestAgentRoles(t *testing.T) {
	fs := testNewFreshservice(t)
	if fs == nil {
		return
	}

	itcnt := 0
	err := fs.IterAgentRoles(ctxbg, nil, func(ar *AgentRole) error {
		itcnt++
		tlog.Debugf("Iterate agent role #%d: %s", ar.ID, ar.Name)
		return nil
	})
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	tlog.Infof("Iterate %d agent roles", itcnt)
}
