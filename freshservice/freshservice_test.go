package freshservice

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

func testNewFreshservice(t *testing.T) *Client {
	apikey := os.Getenv("FSV_APIKEY")
	if apikey == "" {
		t.Skip("FSV_APIKEY not set")
		return nil
	}

	domain := os.Getenv("FSV_DOMAIN")
	if domain == "" {
		t.Skip("FSV_DOMAIN not set")
		return nil
	}

	logger := tlog.GetLogger("FDK")

	fsv := &Client{
		Domain:    domain,
		APIKey:    apikey,
		Transport: httplog.LoggingRoundTripper(logger),
		Retryer:   NewRetryer(time.Second*3, 1, logger),
	}

	return fsv
}
