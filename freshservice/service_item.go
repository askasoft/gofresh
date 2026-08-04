package freshservice

import (
	"github.com/askasoft/pango/num"
	"github.com/askasoft/pango/str"
)

type ServiceItemType int
type ServiceItemVisibility int
type GroupVisibility int
type AgentGroupVisibility int

const (
	ServiceItemNormal ServiceItemType = 1
	ServiceItemLoaner ServiceItemType = 2

	ServiceItemDraft     ServiceItemVisibility = 1
	ServiceItemPublished ServiceItemVisibility = 2

	AgentGroupVisibilityAll        AgentGroupVisibility = 1
	AgentGroupVisibilityGroups     AgentGroupVisibility = 2
	AgentGroupVisibilityWorkspaces AgentGroupVisibility = 3

	GroupVisibilityAll        GroupVisibility = 1
	GroupVisibilityRestricted GroupVisibility = 2
)

func (t ServiceItemType) String() string {
	switch t {
	case ServiceItemNormal:
		return "Normal"
	case ServiceItemLoaner:
		return "Loaner"
	default:
		return num.Itoa(int(t))
	}
}

func ParseServiceItemType(s string) ServiceItemType {
	switch str.ToLower(s) {
	case "normal":
		return ServiceItemNormal
	case "loaner":
		return ServiceItemLoaner
	default:
		return 0
	}
}
func (v ServiceItemVisibility) String() string {
	switch v {
	case ServiceItemDraft:
		return "Draft"
	case ServiceItemPublished:
		return "Published"
	default:
		return num.Itoa(int(v))
	}
}

func ParseServiceItemVisibility(s string) ServiceItemVisibility {
	switch str.ToLower(s) {
	case "draft":
		return ServiceItemDraft
	case "published":
		return ServiceItemPublished
	default:
		return 0
	}
}

type AgentGroupVisibilities struct {
	GroupID []int64 `json:"group_id,omitempty"`
}

type AgentWorkspaceVisibilities struct {
	WorkspaceIDs []int64 `json:"workspace_ids,omitempty"`
}

// ServiceItem represents a Freshservice Service Item.
type ServiceItem struct {
	ID int64 `json:"id"`

	// Freshservice / Freshservice for Business Teams: ID of the workspace to which the service item belongs.
	// Freshservice for MSPs: ID of the clients the service item belongs to. For Freshservice for MSP, this value is always set to 1.
	WorkspaceID int64 `json:"workspace_id,omitempty"`

	// Specifies the clients applicable to a given service item.
	// Use [1] to indicate applicability to all clients and [-1] to indicate applicability to no clients.
	// Otherwise, list the specific client IDs.
	ApplicableToWorkspaceIDs []int64 `json:"applicable_to_workspace_ids,omitempty"`

	Name string `json:"name,omitempty"`

	// Description of the service item
	Description string `json:"description,omitempty"`

	// Short Description of the service item
	ShortDescription string `json:"short_description,omitempty"`

	// Unique ID of the service item specific to your account
	DisplayID int64 `json:"display_id,omitempty"`

	// Unique ID of the category of the service item
	CategoryID int64 `json:"category_id,omitempty"`

	// The ID of the product mapped to the item. Returns null if no product is mapped
	ProductID *int64 `json:"product_id,omitempty"`

	// ‘1’ indicates a normal item. ‘2’ indicates a loaner item
	ItemType ServiceItemType `json:"item_type,omitempty"`

	Quantity int `json:"quantity,omitempty"`

	// Set as True to allow the requester to request for more than 1 quantity
	AllowQuantity bool `json:"allow_quantity,omitempty"`

	// Unique id of the asset type associated with the product
	CITypeID *int64 `json:"ci_type_id,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	// ‘1’ denotes draft and ‘2’ denotes published.
	Visibility ServiceItemVisibility `json:"visibility,omitempty"`

	// 1 denotes visibility to all requesters. 2 for restricted visibility
	GroupVisibility GroupVisibility `json:"group_visibility,omitempty"`

	// undocumented
	GroupVisibilitiesGroupID []int64 `json:"group_visibilities_group_id"`

	// undocumented
	GroupVisibilitiesItemID []int64 `json:"group_visibilities_item_id"`

	// This attribute describes the collection of agents who can view this service item in the service catalog. Possible values: 1: All agents. 2: Agents in specific groups. 3: Agents in specific workspaces.	All products
	AgentGroupVisibility AgentGroupVisibility `json:"agent_group_visibility,omitempty"`

	// undocumented
	AgentGroupVisibilitiesGroupID []int64 `json:"agent_group_visibilities_group_id"`

	// This attribute contains the group_id attribute, and is applicable only if agent_group_visibility is set to 2 (Agents in specific groups).
	AgentGroupVisibilities *AgentGroupVisibilities `json:"agent_group_visibilities,omitempty"`

	// This attribute contains the workspace_ids attribute, and is applicable only if agent_group_visibility is set to 3 (Agents in specific workspaces).
	AgentWorkspaceVisibilities *AgentWorkspaceVisibilities `json:"agent_workspace_visibilities,omitempty"`

	// 	Estimated delivery time of the item (in hours)
	DeliveryTime int `json:"delivery_time,omitempty"`

	// Set to True if delivery time of the item should be visible to the requester
	DeliveryTimeVisibility bool `json:"delivery_time_visibility,omitempty"`

	// Cost of the service item
	Cost string `json:"cost,omitempty"`

	// Set to True if cost should be visible to the requester
	CostVisibility bool `json:"cost_visibility,omitempty"`

	// Config indicating the template of the service request subject
	Configs map[string]any `json:"configs,omitempty"`

	// Set to True if item is “bot ready”
	Botified bool `json:"botified,omitempty"`

	// Set to True if requester is allowed to attach a file
	AllowAttachments bool `json:"allow_attachments,omitempty"`

	// Boolean indicating whether the item contains child items
	IsBundle bool `json:"is_bundle,omitempty"`

	// Boolean indicating whether child items will be created as separate service request
	CreateChild bool `json:"create_child,omitempty"`

	// Custom fields associated with the service item
	CustomFields any `json:"custom_fields,omitempty"`

	// Child Service Items attached to this item
	ChildItems []ServiceItem `json:"child_items,omitempty"`

	// Undocumented
	IconName string `json:"icon_name,omitempty"`
	IconURL  string `json:"icon_url,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`
	UpdatedAt Time `json:"updated_at,omitempty"`
}

func (si *ServiceItem) String() string {
	return toString(si)
}

// ServiceItemCreate represents a Freshservice Service Item.
type ServiceItemCreate struct {
	// Freshservice / Freshservice for Business Teams: ID of the workspace to which the service item belongs.
	// Freshservice for MSPs: ID of the clients the service item belongs to. For Freshservice for MSP, this value is always set to 1.
	WorkspaceID int64 `json:"workspace_id,omitempty"`

	// Specifies the clients applicable to a given service item.
	// Use [1] to indicate applicability to all clients and [-1] to indicate applicability to no clients.
	// Otherwise, list the specific client IDs.
	ApplicableToWorkspaceIDs []int64 `json:"applicable_to_workspace_ids,omitempty"`

	Name string `json:"name,omitempty"`

	// Unique ID of the category of the service item
	CategoryID int64 `json:"category_id,omitempty"`

	// Short Description of the service item
	ShortDescription string `json:"short_description,omitempty"`

	// Description of the service item
	Description string `json:"description,omitempty"`

	// ‘1’ denotes draft and ‘2’ denotes published.
	Visibility ServiceItemVisibility `json:"visibility,omitempty"`

	// Cost of the service item
	Cost float64 `json:"cost,omitempty"`

	// Set to True if cost should be visible to the requester
	CostVisibility bool `json:"cost_visibility,omitempty"`

	// 	Estimated delivery time of the item (in hours)
	DeliveryTime int `json:"delivery_time,omitempty"`

	// Set to True if delivery time of the item should be visible to the requester
	DeliveryTimeVisibility bool `json:"delivery_time_visibility,omitempty"`

	// 1 denotes visibility to all requesters. 2 for restricted visibility
	GroupVisibility GroupVisibility `json:"group_visibility,omitempty"`

	// This attribute describes the collection of agents who can view this service item in the service catalog. Possible values: 1: All agents. 2: Agents in specific groups. 3: Agents in specific workspaces.	All products
	AgentGroupVisibility AgentGroupVisibility `json:"agent_group_visibility,omitempty"`

	// This attribute contains the group_id attribute, and is applicable only if agent_group_visibility is set to 2 (Agents in specific groups).
	AgentGroupVisibilities *AgentGroupVisibilities `json:"agent_group_visibilities,omitempty"`

	// This attribute contains the workspace_ids attribute, and is applicable only if agent_group_visibility is set to 3 (Agents in specific workspaces).
	AgentWorkspaceVisibilities *AgentWorkspaceVisibilities `json:"agent_workspace_visibilities,omitempty"`

	// This attribute contains the array of agent group identifiers which have been granted access to the service item.
	GroupID []int64 `json:"group_id,omitempty"`

	// This attribute contains the array of workspace identifiers which have been granted access to the service item.
	WorkspaceIDs []int64 `json:"workspace_ids,omitempty"`

	// Custom fields associated with the service item.
	CustomFields []map[string]any `json:"custom_fields,omitempty"`

	// Name of the custom field.
	Label string `json:"label,omitempty"`

	// Placeholder value for the custom field.	All products
	Placeholder string `json:"placeholder,omitempty"`

	// Indicates whether the Field is required or not during the form submission.
	Required bool `json:"required,omitempty"`

	// Contains Options available for selection of the field.
	FieldOptions map[string]any `json:"field_options,omitempty"`

	// Requester can view the field. This field should be set to True if we want requester_can_edit to be Set as True
	DisplayToRequester bool `json:"displayed_to_requester,omitempty"`

	// Requester can edit the field.
	RequesterCanEdit bool `json:"requester_can_edit,omitempty"`

	// 	Dropdown choices for a custom field.
	Choices []map[string]any `json:"choices,omitempty"`

	// Contains dropdown choice values.
	Value string `json:"value,omitempty"`

	// Denotes the Custom field Type.
	FieldType string `json:"field_type,omitempty"`
}

func (sic *ServiceItemCreate) String() string {
	return toString(sic)
}

type ServiceItemUpdate = ServiceItemCreate

type serviceItemResult struct {
	ServiceItem *ServiceItem `json:"service_item,omitempty"`
}

type serviceItemsResult struct {
	ServiceItems []*ServiceItem `json:"service_items,omitempty"`
}
