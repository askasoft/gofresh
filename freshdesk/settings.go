package freshdesk

import "context"

type HelpdeskSettings struct {
	PrimaryLanguage string `json:"primary_language,omitempty"`

	SupportedLanguages []string `json:"supported_languages,omitempty"`

	PortalLanguages []string `json:"portal_languages,omitempty"`
}

func (hs *HelpdeskSettings) String() string {
	return toString(hs)
}

func (c *Client) GetHelpdeskSettings(ctx context.Context) (*HelpdeskSettings, error) {
	url := c.Endpoint("/settings/helpdesk")
	hs := &HelpdeskSettings{}
	err := c.DoGet(ctx, url, hs)
	return hs, err
}
