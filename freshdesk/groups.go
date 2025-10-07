package freshdesk

import "context"

// ---------------------------------------------------
// Group

type ListGroupsOption = PageOption

func (c *Client) GetGroup(ctx context.Context, gid int64) (*Group, error) {
	url := c.Endpoint("/groups/%d", gid)
	group := &Group{}
	err := c.DoGet(ctx, url, group)
	return group, err
}

func (c *Client) CreateGroup(ctx context.Context, group *GroupCreate) (*Group, error) {
	url := c.Endpoint("/groups")
	result := &Group{}
	if err := c.DoPost(ctx, url, group, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ListGroups(ctx context.Context, lgo *ListGroupsOption) ([]*Group, bool, error) {
	url := c.Endpoint("/groups")
	groups := []*Group{}
	next, err := c.DoList(ctx, url, lgo, &groups)
	return groups, next, err
}

func (c *Client) IterGroups(ctx context.Context, lgo *ListGroupsOption, igf func(*Group) error) error {
	if lgo == nil {
		lgo = &ListGroupsOption{}
	}
	if lgo.Page < 1 {
		lgo.Page = 1
	}
	if lgo.PerPage < 1 {
		lgo.PerPage = 100
	}

	for {
		groups, next, err := c.ListGroups(ctx, lgo)
		if err != nil {
			return err
		}
		for _, g := range groups {
			if err = igf(g); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lgo.Page++
	}
	return nil
}

func (c *Client) UpdateGroup(ctx context.Context, gid int64, group *GroupUpdate) (*Group, error) {
	url := c.Endpoint("/groups/%d", gid)
	result := &Group{}
	if err := c.DoPut(ctx, url, group, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteGroup(ctx context.Context, gid int64) error {
	url := c.Endpoint("/groups/%d", gid)
	return c.DoDelete(ctx, url)
}
