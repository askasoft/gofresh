package freshservice

import "context"

// ---------------------------------------------------
// Workspace

type ListWorkspacesOption = PageOption

func (c *Client) GetWorkspace(ctx context.Context, id int64) (*Workspace, error) {
	url := c.Endpoint("/workspaces/%d", id)
	result := &workspaceResult{}
	err := c.DoGet(ctx, url, result)
	return result.Workspace, err
}

func (c *Client) ListWorkspaces(ctx context.Context, lwo *ListWorkspacesOption) ([]*Workspace, bool, error) {
	url := c.Endpoint("/workspaces")
	result := &workspacesResult{}
	next, err := c.DoList(ctx, url, lwo, result)
	return result.Workspaces, next, err
}

func (c *Client) IterWorkspaces(ctx context.Context, lwo *ListWorkspacesOption, iwf func(*Workspace) error) error {
	if lwo == nil {
		lwo = &ListWorkspacesOption{}
	}
	if lwo.Page < 1 {
		lwo.Page = 1
	}
	if lwo.PerPage < 1 {
		lwo.PerPage = 100
	}

	for {
		ws, next, err := c.ListWorkspaces(ctx, lwo)
		if err != nil {
			return err
		}
		for _, ag := range ws {
			if err = iwf(ag); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lwo.Page++
	}
	return nil
}
