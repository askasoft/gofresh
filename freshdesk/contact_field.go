package freshdesk

type ContactField struct {
	ID int64 `json:"id,omitempty"`

	// Name of the contact field.
	Name string `json:"name,omitempty"`

	// Label of the field for display
	Label string `json:"label,omitempty"`

	// Position in which the contact field is displayed in the form
	Position int `json:"field_type,omitempty"`

	// True if the field is a not a custom field.
	Default bool `json:"default,omitempty"`

	// For custom contact fields, The type of value associated with the field will be given (Examples custom_date, custom_text...)
	Type string `json:"type,omitempty"`

	// EditableInSignup Set to true if the field can be updated by customers during signup. The default Value is false
	EditableInSignup bool `json:"editable_in_signup,omitempty"`

	// Set to true if the field can be updated by customers
	CustomersCanEdit bool `json:"customers_can_edit,omitempty"`

	// Display name for the field (as seen in the customer portal)
	LabelForCustomers string `json:"label_for_customers,omitempty"`

	// Set to true if the field is mandatory for Agents
	RequiredForAgents bool `json:"required_for_agents,omitempty"`

	// Set to true if the field is mandatory in the customer portal
	RequiredForCustomers bool `json:"required_for_customer,omitempty"`

	// Set to true if the field is displayed in the customer portal
	DisplayedToCustomers bool `json:"displayed_to_customers,omitempty"`

	// List of values supported by the field.
	Choices any `json:"choices,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`

	UpdatedAt Time `json:"updated_at,omitempty"`
}

func (cf *ContactField) String() string {
	return toString(cf)
}

type ContactFieldCreate struct {
	// Label of the field for display
	Label string `json:"label,omitempty"`

	// Display name for the field (as seen in the customer portal)
	LabelForCustomers string `json:"label_for_customers,omitempty"`

	// For custom ticket fields, The type of value associated with the field will be given (Examples custom_date, custom_text...)
	Type string `json:"type,omitempty"`

	// Position in which the ticket field is displayed in the form
	Position int `json:"field_type,omitempty"`

	// EditableInSignup Set to true if the field can be updated by customers during signup. The default Value is false
	EditableInSignup bool `json:"editable_in_signup,omitempty"`

	// Set to true if the field is mandatory for Agents
	RequiredForAgents bool `json:"required_for_agents,omitempty"`

	// Set to true if the field can be updated by customers
	CustomersCanEdit bool `json:"customers_can_edit,omitempty"`

	// Set to true if the field is mandatory in the customer portal
	RequiredForCustomers bool `json:"required_for_customer,omitempty"`

	// Set to true if the field is displayed in the customer portal
	DisplayedForCustomers bool `json:"displayed_for_customers,omitempty"`

	// List of values supported by the field.
	Choices any `json:"choices,omitempty"`
}

func (cf *ContactFieldCreate) String() string {
	return toString(cf)
}

type ContactFieldUpdate = ContactFieldCreate
