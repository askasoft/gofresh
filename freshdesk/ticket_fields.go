package freshdesk

import (
	"context"
	"strings"
)

// ---------------------------------------------------
// Ticket Fields

const (
	TicketFieldIncludeSection = "section"
)

// List All Ticket Fields
func (c *Client) ListTicketFields(ctx context.Context, types ...string) ([]*TicketField, error) {
	url := c.Endpoint("/admin/ticket_fields")
	if len(types) > 0 {
		s := strings.Join(types, ",")
		url += "?type=" + s
	}

	fields := []*TicketField{}
	err := c.DoGet(ctx, url, &fields)
	return fields, err
}

func (c *Client) CreateTicketField(ctx context.Context, tf *TicketFieldCreate) (*TicketField, error) {
	url := c.Endpoint("/admin/ticket_fields")
	result := &TicketField{}
	if err := c.DoPost(ctx, url, tf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTicketField View a Ticket Field
// include: conversations, requester, company, stats
func (c *Client) GetTicketField(ctx context.Context, fid int64, include ...string) (*TicketField, error) {
	url := c.Endpoint("/admin/ticket_fields/%d", fid)
	if len(include) > 0 {
		s := strings.Join(include, ",")
		url += "?include=" + s
	}

	result := &TicketField{}
	err := c.DoGet(ctx, url, result)
	return result, err
}

func (c *Client) UpdateTicketField(ctx context.Context, fid int64, field *TicketFieldUpdate) (*TicketField, error) {
	url := c.Endpoint("/admin/ticket_fields/%d", fid)
	result := &TicketField{}
	if err := c.DoPut(ctx, url, field, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteTicketField(ctx context.Context, fid int64) error {
	url := c.Endpoint("/admin/ticket_fields/%d", fid)
	return c.DoDelete(ctx, url)
}
