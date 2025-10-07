package freshdesk

import "context"

// ---------------------------------------------------
// Automation

type ListAutomationRulesOption = PageOption

func (c *Client) ListAutomationRules(ctx context.Context, aType AutomationType, laro *ListAutomationRulesOption) ([]*AutomationRule, bool, error) {
	url := c.Endpoint("/automations/%d/rules", aType)
	rules := []*AutomationRule{}
	next, err := c.DoList(ctx, url, laro, &rules)
	return rules, next, err
}

func (c *Client) IterAutomationRules(ctx context.Context, aType AutomationType, laro *ListAutomationRulesOption, iarf func(*AutomationRule) error) error {
	if laro == nil {
		laro = &ListAutomationRulesOption{}
	}
	if laro.Page < 1 {
		laro.Page = 1
	}
	if laro.PerPage < 1 {
		laro.PerPage = 100
	}

	for {
		ars, next, err := c.ListAutomationRules(ctx, aType, laro)
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

func (c *Client) GetAutomationRule(ctx context.Context, aType AutomationType, rid int64) (*AutomationRule, error) {
	url := c.Endpoint("/automations/%d/rules/%d", aType, rid)
	rule := &AutomationRule{}
	err := c.DoGet(ctx, url, rule)
	return rule, err
}

func (c *Client) DeleteAutomationRule(ctx context.Context, aType AutomationType, rid int64) error {
	url := c.Endpoint("/automations/%d/rules/%d", aType, rid)
	return c.DoDelete(ctx, url)
}

func (c *Client) CreateAutomationRule(ctx context.Context, aType AutomationType, rule *AutomationRuleCreate) (*AutomationRule, error) {
	url := c.Endpoint("/automations/%d/rules", aType)
	result := &AutomationRule{}
	if err := c.DoPost(ctx, url, rule, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateAutomationRule(ctx context.Context, aType AutomationType, rid int64, rule *AutomationRuleUpdate) (*AutomationRule, error) {
	url := c.Endpoint("/automations/%d/rules/%d", aType, rid)
	result := &AutomationRule{}
	if err := c.DoPut(ctx, url, rule, result); err != nil {
		return nil, err
	}
	return result, nil
}
