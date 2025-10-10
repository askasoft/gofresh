package freshdesk

import (
	"context"
)

// GetJob get job detail
func (c *Client) GetJob(ctx context.Context, jid string) (*Job, error) {
	url := c.Endpoint("/jobs/%s", jid)
	job := &Job{}
	err := c.DoGet(ctx, url, job)
	return job, err
}
