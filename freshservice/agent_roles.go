package freshservice

import "context"

// ---------------------------------------------------
// Agent Role

type ListAgentRolesOption = PageOption

func (c *Client) GetAgentRole(ctx context.Context, id int64) (*AgentRole, error) {
	url := c.Endpoint("/roles/%d", id)
	result := &agentRoleResult{}
	err := c.DoGet(ctx, url, result)
	return result.Role, err
}

func (c *Client) ListAgentRoles(ctx context.Context, laro *ListAgentRolesOption) ([]*AgentRole, bool, error) {
	url := c.Endpoint("/roles")
	result := &agentRolesResult{}
	next, err := c.DoList(ctx, url, laro, result)
	return result.Roles, next, err
}

func (c *Client) IterAgentRoles(ctx context.Context, laro *ListAgentRolesOption, iarf func(*AgentRole) error) error {
	if laro == nil {
		laro = &ListAgentRolesOption{}
	}
	if laro.Page < 1 {
		laro.Page = 1
	}
	if laro.PerPage < 1 {
		laro.PerPage = 100
	}

	for {
		ars, next, err := c.ListAgentRoles(ctx, laro)
		if err != nil {
			return err
		}
		for _, ar := range ars {
			if err = iarf(ar); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		laro.Page++
	}
	return nil
}
