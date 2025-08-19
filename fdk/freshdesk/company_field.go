package freshdesk

type CompanyField struct {
	ID int64 `json:"id,omitempty"`

	// Name of the company field.
	Name string `json:"name,omitempty"`

	// Label of the field for display
	Label string `json:"label,omitempty"`

	// For custom conmpanyfields, The type of value associated with the field will be given (Examples custom_date, custom_text...)
	Type string `json:"type,omitempty"`

	// Position in which the company field is displayed in the form
	Position int `json:"field_type,omitempty"`

	// True if the field is a not a custom field.
	Default bool `json:"default,omitempty"`

	// Set to true if the field is mandatory for Agents
	RequiredForAgents bool `json:"required_for_agents,omitempty"`

	// List of values supported by the field.
	Choices any `json:"choices,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`

	UpdatedAt Time `json:"updated_at,omitempty"`
}

func (cf *CompanyField) String() string {
	return toString(cf)
}

type CompanyFieldCreate struct {
	// Label of the field for display
	Label string `json:"label,omitempty"`

	// For custom conmpanyfields, The type of value associated with the field will be given (Examples custom_date, custom_text...)
	Type string `json:"type,omitempty"`

	// Position in which the company field is displayed in the form
	Position int `json:"field_type,omitempty"`

	// Set to true if the field is mandatory for Agents
	RequiredForAgents bool `json:"required_for_agents,omitempty"`

	// List of values supported by the field.
	Choices any `json:"choices,omitempty"`
}

func (cf *CompanyFieldCreate) String() string {
	return toString(cf)
}

type CompanyFieldUpdate = CompanyFieldCreate
