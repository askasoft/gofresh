package freshdesk

import (
	"testing"
)

func TestGetHelpdeskSettings(t *testing.T) {
	fd := testNewFreshdesk(t)
	if fd == nil {
		return
	}

	hs, err := fd.GetHelpdeskSettings(ctxbg)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	tlog.Debug(hs)
}
