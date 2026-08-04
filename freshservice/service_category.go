package freshservice

// ServiceCategory represents a Freshservice Service Category.
type ServiceCategory struct {
	ID int64 `json:"id"`

	// Freshservice / Freshservice for Business Teams: ID of the workspace to which the service item belongs.
	// Freshservice for MSPs: ID of the clients the service item belongs to. For Freshservice for MSP, this value is always set to 1.
	WorkspaceID int64 `json:"workspace_id,omitempty"`

	Name string `json:"name,omitempty"`

	// Description of the service category
	Description string `json:"description,omitempty"`

	// Number denoting the position of category in service catalog
	Position int `json:"position,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`
	UpdatedAt Time `json:"updated_at,omitempty"`
}

func (sc *ServiceCategory) String() string {
	return toString(sc)
}

type serviceCategoriesResult struct {
	ServiceCategories []*ServiceCategory `json:"service_categories,omitempty"`
}
