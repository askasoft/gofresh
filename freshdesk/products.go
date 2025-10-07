package freshdesk

import "context"

// ---------------------------------------------------
// Product

type ListProductsOption = PageOption

func (c *Client) GetProduct(ctx context.Context, id int64) (*Product, error) {
	url := c.Endpoint("/products/%d", id)
	product := &Product{}
	err := c.DoGet(ctx, url, product)
	return product, err
}

func (c *Client) ListProducts(ctx context.Context, lpo *ListProductsOption) ([]*Product, bool, error) {
	url := c.Endpoint("/products")
	products := []*Product{}
	next, err := c.DoList(ctx, url, lpo, &products)
	return products, next, err
}

func (c *Client) IterProducts(ctx context.Context, lpo *ListProductsOption, ipf func(*Product) error) error {
	if lpo == nil {
		lpo = &ListProductsOption{}
	}
	if lpo.Page < 1 {
		lpo.Page = 1
	}
	if lpo.PerPage < 1 {
		lpo.PerPage = 100
	}

	for {
		ps, next, err := c.ListProducts(ctx, lpo)
		if err != nil {
			return err
		}
		for _, ag := range ps {
			if err = ipf(ag); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lpo.Page++
	}
	return nil
}
