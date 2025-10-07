package freshservice

import "context"

// ---------------------------------------------------
// Solutions

// PerPage: 1 ~ 100, default: 30
type ListCategoriesOption = PageOption

// PerPage: 1 ~ 100, default: 30
type ListFoldersOption struct {
	categoryID int64
	Page       int
	PerPage    int
}

func (lfo *ListFoldersOption) IsNil() bool {
	return lfo == nil
}

func (lfo *ListFoldersOption) Values() Values {
	q := Values{}
	q.SetInt64("category_id", lfo.categoryID)
	q.SetInt("page", lfo.Page)
	q.SetInt("per_page", lfo.PerPage)
	return q
}

// PerPage: 1 ~ 100, default: 30
type ListArticlesOption struct {
	folderID int64
	Page     int
	PerPage  int
}

func (lao *ListArticlesOption) IsNil() bool {
	return lao == nil
}

func (lao *ListArticlesOption) Values() Values {
	q := Values{}
	q.SetInt64("folder_id", lao.folderID)
	q.SetInt("page", lao.Page)
	q.SetInt("per_page", lao.PerPage)
	return q
}

type SearchArticlesOption struct {
	SearchTerm string // The keywords for which the solution articles have to be searched.
	UserEmail  string // By default, the API will search the articles for the user whose API key is provided. If you want to search articles for a different user, please provide their user_email.
	Page       int
	PerPage    int
}

func (sao *SearchArticlesOption) IsNil() bool {
	return sao == nil
}

func (sao *SearchArticlesOption) Values() Values {
	q := Values{}
	q.SetString("search_term", sao.SearchTerm)
	q.SetString("user_email", sao.UserEmail)
	q.SetInt("page", sao.Page)
	q.SetInt("per_page", sao.PerPage)
	return q
}

func (c *Client) CreateCategory(ctx context.Context, category *CategoryCreate) (*Category, error) {
	url := c.Endpoint("/solutions/categories")
	result := &categoryResult{}
	if err := c.DoPost(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result.Category, nil
}

func (c *Client) UpdateCategory(ctx context.Context, cid int64, category *CategoryUpdate) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d", cid)
	result := &categoryResult{}
	if err := c.DoPut(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result.Category, nil
}

func (c *Client) GetCategory(ctx context.Context, cid int64) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d", cid)
	result := &categoryResult{}
	err := c.DoGet(ctx, url, result)
	return result.Category, err
}

func (c *Client) ListCategories(ctx context.Context, lco *ListCategoriesOption) ([]*Category, bool, error) {
	url := c.Endpoint("/solutions/categories")
	result := &categoriesResult{}
	next, err := c.DoList(ctx, url, lco, result)
	return result.Categories, next, err
}

func (c *Client) IterCategories(ctx context.Context, lco *ListCategoriesOption, icf func(*Category) error) error {
	if lco == nil {
		lco = &ListCategoriesOption{}
	}
	if lco.Page < 1 {
		lco.Page = 1
	}
	if lco.PerPage < 1 {
		lco.PerPage = 100
	}

	for {
		categories, next, err := c.ListCategories(ctx, lco)
		if err != nil {
			return err
		}
		for _, c := range categories {
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

func (c *Client) DeleteCategory(ctx context.Context, cid int64) error {
	url := c.Endpoint("/solutions/categories/%d", cid)
	return c.DoDelete(ctx, url)
}

func (c *Client) CreateFolder(ctx context.Context, folder *FolderCreate) (*Folder, error) {
	url := c.Endpoint("/solutions/folders")
	result := &folderResult{}
	if err := c.DoPost(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result.Foler, nil
}

func (c *Client) UpdateFolder(ctx context.Context, fid int64, folder *FolderUpdate) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d", fid)
	result := &folderResult{}
	if err := c.DoPut(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result.Foler, nil
}

func (c *Client) GetFolder(ctx context.Context, fid int64) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d", fid)
	result := &folderResult{}
	err := c.DoGet(ctx, url, result)
	return result.Foler, err
}

func (c *Client) ListCategoryFolders(ctx context.Context, cid int64, lfo *ListFoldersOption) ([]*Folder, bool, error) {
	if lfo == nil {
		lfo = &ListFoldersOption{}
	}
	lfo.categoryID = cid

	url := c.Endpoint("/solutions/folders")
	result := &foldersResult{}
	next, err := c.DoList(ctx, url, lfo, result)
	return result.Folders, next, err
}

func (c *Client) IterCategoryFolders(ctx context.Context, cid int64, lfo *ListFoldersOption, iff func(*Folder) error) error {
	if lfo == nil {
		lfo = &ListFoldersOption{}
	}
	if lfo.Page < 1 {
		lfo.Page = 1
	}
	if lfo.PerPage < 1 {
		lfo.PerPage = 100
	}

	for {
		folders, next, err := c.ListCategoryFolders(ctx, cid, lfo)
		if err != nil {
			return err
		}
		for _, f := range folders {
			if err = iff(f); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lfo.Page++
	}
	return nil
}

func (c *Client) DeleteFolder(ctx context.Context, fid int64) error {
	url := c.Endpoint("/solutions/folders/%d", fid)
	return c.DoDelete(ctx, url)
}

func (c *Client) CreateArticle(ctx context.Context, article *ArticleCreate) (*Article, error) {
	url := c.Endpoint("/solutions/articles")
	result := &articleResult{}
	if err := c.DoPost(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result.Article, nil
}

func (c *Client) SendArticleToApproval(ctx context.Context, aid int64) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d/send_for_approval", aid)
	result := &articleResult{}
	if err := c.DoPut(ctx, url, nil, result); err != nil {
		return nil, err
	}
	return result.Article, nil
}

func (c *Client) UpdateArticle(ctx context.Context, aid int64, article *ArticleUpdate) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d", aid)
	result := &articleResult{}
	if err := c.DoPut(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result.Article, nil
}

func (c *Client) GetArticle(ctx context.Context, aid int64) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d", aid)
	result := &articleResult{}
	err := c.DoGet(ctx, url, result)
	return result.Article, err
}

func (c *Client) ListFolderArticles(ctx context.Context, fid int64, lao *ListArticlesOption) ([]*ArticleInfo, bool, error) {
	if lao == nil {
		lao = &ListArticlesOption{}
	}
	lao.folderID = fid

	url := c.Endpoint("/solutions/articles")
	result := &articlesResult{}
	next, err := c.DoList(ctx, url, lao, result)
	for _, ai := range result.Articles {
		ai.normalize()
	}
	return result.Articles, next, err
}

func (c *Client) IterFolderArticles(ctx context.Context, fid int64, lao *ListArticlesOption, iaf func(*ArticleInfo) error) error {
	if lao == nil {
		lao = &ListArticlesOption{}
	}
	if lao.Page < 1 {
		lao.Page = 1
	}
	if lao.PerPage < 1 {
		lao.PerPage = 100
	}

	for {
		articles, next, err := c.ListFolderArticles(ctx, fid, lao)
		if err != nil {
			return err
		}
		for _, a := range articles {
			if err = iaf(a); err != nil {
				return err
			}
		}
		if !next {
			break
		}
		lao.Page++
	}
	return nil
}

func (c *Client) DeleteArticle(ctx context.Context, aid int64) error {
	url := c.Endpoint("/solutions/articles/%d", aid)
	return c.DoDelete(ctx, url)
}

func (c *Client) SearchArticles(ctx context.Context, sao *SearchArticlesOption) ([]*ArticleInfo, bool, error) {
	url := c.Endpoint("/solutions/articles/search")
	result := &articlesResult{}
	next, err := c.DoList(ctx, url, sao, result)
	for _, ai := range result.Articles {
		ai.normalize()
	}
	return result.Articles, next, err
}

// func (c *Client) DeleteArticleAttachment(aid, tid int64) error {
// 	url := c.Endpoint("/solutions/articles/%d/attachments/%d", aid, tid)
// 	return c.DoDelete(ctx, url)
// }
