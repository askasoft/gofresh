package freshdesk

import "context"

// ---------------------------------------------------
// Role

type ListRolesOption = PageOption

func (c *Client) GetRole(ctx context.Context, rid int64) (*Role, error) {
	url := c.Endpoint("/roles/%d", rid)
	role := &Role{}
	err := c.DoGet(ctx, url, role)
	return role, err
}

func (c *Client) ListRoles(ctx context.Context, lro *ListRolesOption) ([]*Role, bool, error) {
	url := c.Endpoint("/roles")
	roles := []*Role{}
	next, err := c.DoList(ctx, url, lro, &roles)
	return roles, next, err
}

func (c *Client) IterRoles(ctx context.Context, lro *ListRolesOption, irf func(*Role) error) error {
	if lro == nil {
		lro = &ListRolesOption{}
	}
	if lro.Page < 1 {
		lro.Page = 1
	}
	if lro.PerPage < 1 {
		lro.PerPage = 100
	}

	for {
		roles, next, err := c.ListRoles(ctx, lro)
		if err != nil {
			return err
		}
		for _, g := range roles {
			if err = irf(g); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lro.Page++
	}
	return nil
}
