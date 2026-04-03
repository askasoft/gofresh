package freshdesk

import (
	"os"
	"testing"
	"time"

	"github.com/askasoft/pango/log"
	"github.com/askasoft/pango/log/httplog"
)

var tlog *log.Log

func init() {
	tlog = log.NewLog()
	tlog.SetLevel(log.LevelInfo)
}

func testNewFreshdesk(t *testing.T) *Client {
	apikey := os.Getenv("FDK_APIKEY")
	if apikey == "" {
		t.Skip("FDK_APIKEY not set")
		return nil
	}

	domain := os.Getenv("FDK_DOMAIN")
	if domain == "" {
		t.Skip("FDK_DOMAIN not set")
		return nil
	}

	logger := tlog.GetLogger("FDK")

	fdk := &Client{
		Domain:    domain,
		APIKey:    apikey,
		Transport: httplog.LoggingRoundTripper(logger),
		Retryer:   NewRetryer(time.Second*3, 1, logger),
	}

	return fdk
}
