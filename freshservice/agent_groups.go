package freshservice

import "context"

// ---------------------------------------------------
// Agent Group

type ListAgentGroupsOption = PageOption

func (c *Client) CreateAgentGroup(ctx context.Context, ag *AgentGroupCreate) (*AgentGroup, error) {
	url := c.Endpoint("/groups")
	result := &agentGroupResult{}
	if err := c.DoPost(ctx, url, ag, result); err != nil {
		return nil, err
	}
	return result.Group, nil
}

func (c *Client) GetAgentGroup(ctx context.Context, id int64) (*AgentGroup, error) {
	url := c.Endpoint("/groups/%d", id)
	result := &agentGroupResult{}
	err := c.DoGet(ctx, url, result)
	return result.Group, err
}

func (c *Client) ListAgentGroups(ctx context.Context, lago *ListAgentGroupsOption) ([]*AgentGroup, bool, error) {
	url := c.Endpoint("/groups")
	result := &agentGroupsResult{}
	next, err := c.DoList(ctx, url, lago, result)
	return result.Groups, next, err
}

func (c *Client) IterAgentGroups(ctx context.Context, lago *ListAgentGroupsOption, iagf func(*AgentGroup) error) error {
	if lago == nil {
		lago = &ListAgentRolesOption{}
	}
	if lago.Page < 1 {
		lago.Page = 1
	}
	if lago.PerPage < 1 {
		lago.PerPage = 100
	}

	for {
		ags, next, err := c.ListAgentGroups(ctx, lago)
		if err != nil {
			return err
		}
		for _, ag := range ags {
			if err = iagf(ag); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lago.Page++
	}
	return nil
}

func (c *Client) UpdateAgentGroup(ctx context.Context, id int64, ag *AgentGroupUpdate) (*AgentGroup, error) {
	url := c.Endpoint("/groups/%d", id)
	result := &agentGroupResult{}
	if err := c.DoPut(ctx, url, ag, result); err != nil {
		return nil, err
	}
	return result.Group, nil
}

func (c *Client) DeleteAgentGroup(ctx context.Context, id int64) error {
	url := c.Endpoint("/groups/%d", id)
	return c.DoDelete(ctx, url)
}
