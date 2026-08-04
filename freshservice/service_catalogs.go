package freshservice

import "context"

func (c *Client) GetServiceItem(ctx context.Context, displayID int64) (*ServiceItem, error) {
	url := c.Endpoint("/service_catalog/items/%d", displayID)
	result := &serviceItemResult{}
	err := c.DoGet(ctx, url, result)
	return result.ServiceItem, err
}

type ListServiceItemsOption struct {
	CategoryID              int64 // Filter by category ID
	WorkspaceID             int64 // Filter by workspace ID (defaults to primary workspace if not specified)
	ApplicableToWorkspaceID int64 // Filter by applicable to workspace ID (allowed values: 1, -1 and valid client's ID)
	Page                    int   // Page number for pagination (default: 1)
	PerPage                 int   // Number of items per page (default: 30, max: 30)
}

func (lio *ListServiceItemsOption) IsNil() bool {
	return lio == nil
}

func (lio *ListServiceItemsOption) Values() Values {
	q := Values{}
	q.SetInt64("category_id", lio.CategoryID)
	q.SetInt64("workspace_id", lio.WorkspaceID)
	q.SetInt64("applicable_to_workspace_id", lio.ApplicableToWorkspaceID)
	q.SetInt("page", lio.Page)
	q.SetInt("per_page", lio.PerPage)
	return q
}

func (c *Client) ListServiceItems(ctx context.Context, lio *ListServiceItemsOption) ([]*ServiceItem, bool, error) {
	url := c.Endpoint("/service_catalog/items")
	result := &serviceItemsResult{}
	next, err := c.DoList(ctx, url, lio, result)
	return result.ServiceItems, next, err
}

func (c *Client) IterServiceItems(ctx context.Context, lio *ListServiceItemsOption, iif func(*ServiceItem) error) error {
	if lio == nil {
		lio = &ListServiceItemsOption{}
	}
	if lio.Page < 1 {
		lio.Page = 1
	}
	if lio.PerPage < 1 {
		lio.PerPage = 30
	}

	for {
		items, next, err := c.ListServiceItems(ctx, lio)
		if err != nil {
			return err
		}
		for _, item := range items {
			if err = iif(item); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lio.Page++
	}
	return nil
}

type SearchServiceItemsOption struct {
	SearchItem string // Search term to filter items
	UserEmail  string // Email of user to search on behalf of
	Page       int    // Page number for pagination (default: 1)
	PerPage    int    // Number of items per page (default: 30, max: 30)
}

func (sio *SearchServiceItemsOption) IsNil() bool {
	return sio == nil
}

func (sio *SearchServiceItemsOption) Values() Values {
	q := Values{}
	q.SetString("search_term", sio.SearchItem)
	q.SetString("user_email", sio.UserEmail)
	q.SetInt("page", sio.Page)
	q.SetInt("per_page", sio.PerPage)
	return q
}

func (c *Client) SearchServiceItems(ctx context.Context, sio *SearchServiceItemsOption) ([]*ServiceItem, bool, error) {
	url := c.Endpoint("/service_catalog/items/search")
	result := &serviceItemsResult{}
	next, err := c.DoList(ctx, url, sio, result)
	return result.ServiceItems, next, err
}

func (c *Client) CreateServiceItem(ctx context.Context, item *ServiceItemCreate) (*ServiceItem, error) {
	url := c.Endpoint("/service-catalog/items")
	result := &serviceItemResult{}
	if err := c.DoPost(ctx, url, item, result); err != nil {
		return nil, err
	}
	return result.ServiceItem, nil
}

func (c *Client) UpdateServiceItem(ctx context.Context, displayID int64, item *ServiceItemUpdate) (*ServiceItem, error) {
	url := c.Endpoint("/service-catalog/items/%d", displayID)
	result := &serviceItemResult{}
	if err := c.DoPut(ctx, url, item, result); err != nil {
		return nil, err
	}
	return result.ServiceItem, nil
}

func (c *Client) DeleteServiceItem(ctx context.Context, displayID int64) error {
	url := c.Endpoint("/service-catalog/items/%d", displayID)
	return c.DoDelete(ctx, url)
}

type ListServiceCategoriesOption struct {
	WorkspaceID int64 // Filter by workspace ID (defaults to primary workspace if not specified)
	Page        int   // Page number for pagination (default: 1)
	PerPage     int   // Number of items per page (default: 30, max: 30)
}

func (lco *ListServiceCategoriesOption) IsNil() bool {
	return lco == nil
}

func (lco *ListServiceCategoriesOption) Values() Values {
	q := Values{}
	q.SetInt64("workspace_id", lco.WorkspaceID)
	q.SetInt("page", lco.Page)
	q.SetInt("per_page", lco.PerPage)
	return q
}

func (c *Client) ListServiceCategories(ctx context.Context, lco *ListServiceCategoriesOption) ([]*ServiceCategory, bool, error) {
	url := c.Endpoint("/service_catalog/categories")
	result := &serviceCategoriesResult{}
	next, err := c.DoList(ctx, url, lco, result)
	return result.ServiceCategories, next, err
}

func (c *Client) IterServiceCategories(ctx context.Context, lco *ListServiceCategoriesOption, icf func(*ServiceCategory) error) error {
	if lco == nil {
		lco = &ListServiceCategoriesOption{}
	}
	if lco.Page < 1 {
		lco.Page = 1
	}
	if lco.PerPage < 1 {
		lco.PerPage = 30
	}

	for {
		scs, next, err := c.ListServiceCategories(ctx, lco)
		if err != nil {
			return err
		}
		for _, c := range scs {
			if err = icf(c); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lco.Page++
	}
	return nil
}
