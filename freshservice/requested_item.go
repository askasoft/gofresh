package freshservice

import (
	"github.com/askasoft/pango/num"
	"github.com/askasoft/pango/str"
)

type RequestedItemStage int

const (
	RequestedItemRequested          RequestedItemStage = 1
	RequestedItemDelivered          RequestedItemStage = 2
	RequestedItemCancelled          RequestedItemStage = 3
	RequestedItemFulfilled          RequestedItemStage = 4
	RequestedItemPartiallyFulfilled RequestedItemStage = 5
)

func (s RequestedItemStage) String() string {
	switch s {
	case RequestedItemRequested:
		return "Requested"
	case RequestedItemDelivered:
		return "Delivered"
	case RequestedItemCancelled:
		return "Cancelled"
	case RequestedItemFulfilled:
		return "Fulfilled"
	case RequestedItemPartiallyFulfilled:
		return "Partially Fulfilled"
	default:
		return num.Itoa(int(s))
	}
}

func ParseRequestedItemStage(s string) RequestedItemStage {
	switch str.ToLower(s) {
	case "requested":
		return RequestedItemRequested
	case "delivered":
		return RequestedItemDelivered
	case "cancelled":
		return RequestedItemCancelled
	case "fulfilled":
		return RequestedItemFulfilled
	case "partially fulfilled":
		return RequestedItemPartiallyFulfilled
	default:
		return 0
	}
}

// RequestedItem represents a requested service item attached to a service requesri.
type RequestedItem struct {
	ID             int64          `json:"id,omitempty"`
	Quantity       int            `json:"quantity,omitempty"`         // Number of units of the item needed by the requester. By default it is 1
	Stage          any            `json:"stage,omitempty"`            // Current stage of the requested item
	Loaned         bool           `json:"loaned,omitempty"`           // Indicated whether the requested item is a loaner item
	CostPerRequest float64        `json:"cost_per_request,omitempty"` // Cost of the requested service item
	Remarks        string         `json:"remarks,omitempty"`          // Remarks related to a requested item
	DeliveryTime   int            `json:"delivery_time,omitempty"`    // Estimated delivery time (in hrs)
	IsParent       bool           `json:"is_parent,omitempty"`        // Boolean indicating whether this is the parent service item
	ServiceItemID  int64          `json:"service_item_id,omitempty"`  // Display id of service item unique to your account
	CustomFields   map[string]any `json:"custom_fields,omitempty"`    // Custom fields
	Attachments    []*Attachment  `json:"attachments,omitempty"`      // Attachments of requested service item

	ItemID            int64        `json:"item_id,omitempty"`
	Item              *ServiceItem `json:"item,omitempty"`
	Cost              string       `json:"cost,omitempty"`
	FromDate          Time         `json:"from_date,omitzero"`
	ToDate            Time         `json:"to_date,omitzero"`
	Fulfilled         bool         `json:"fulfilled,omitempty"`
	DocumentFulfilled bool         `json:"document_fulfilled,omitempty"`
	FulfilledQuantity int          `json:"fulfilled_quantity,omitempty"`

	CreatedAt Time `json:"created_at,omitzero"`
	UpdatedAt Time `json:"updated_at,omitzero"`
}

func (ri *RequestedItem) String() string {
	return toString(ri)
}

func (ri *RequestedItem) Files() Files {
	return ((Attachments)(ri.Attachments)).Files()
}

func (ri *RequestedItem) Values() Values {
	vs := Values{}

	vs.SetInt("quantity", ri.Quantity)
	if s, ok := ri.Stage.(RequestedItemStage); ok {
		vs.SetInt("stage", int(s))
	}
	vs.SetBool("loaned", ri.Loaned)
	vs.SetFloat64("cost_per_request", ri.CostPerRequest)
	vs.SetString("remarks", ri.Remarks)
	vs.SetInt("delivery_time", ri.DeliveryTime)
	vs.SetBool("is_parent", ri.IsParent)
	vs.SetMap("custom_fields", ri.CustomFields)

	return vs
}

type RequestedItemCreate = RequestedItem
type RequestedItemUpdate = RequestedItemCreate

type requestedItemResult struct {
	RequestedItem *RequestedItem `json:"requested_item,omitempty"`
}

type requestedItemsResult struct {
	RequestedItems []*RequestedItem `json:"requested_items,omitempty"`
}
