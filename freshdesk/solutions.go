package freshdesk

import (
	"context"
	"net/url"
)

// ---------------------------------------------------
// Solutions

// PerPage: 1 ~ 100, default: 30
type ListCategoriesOption = PageOption

// PerPage: 1 ~ 100, default: 30
type ListFoldersOption = PageOption

// PerPage: 1 ~ 100, default: 30
type ListArticlesOption = PageOption

func (c *Client) CreateCategory(ctx context.Context, category *CategoryCreate) (*Category, error) {
	url := c.Endpoint("/solutions/categories")
	result := &Category{}
	if err := c.DoPost(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateCategoryTranslated(ctx context.Context, cid int64, lang string, category *CategoryCreate) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d/%s", cid, lang)
	result := &Category{}
	if err := c.DoPost(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateCategory(ctx context.Context, cid int64, category *CategoryUpdate) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d", cid)
	result := &Category{}
	if err := c.DoPut(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateCategoryTranslated(ctx context.Context, cid int64, lang string, category *CategoryUpdate) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d/%s", cid, lang)
	result := &Category{}
	if err := c.DoPut(ctx, url, category, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetCategory(ctx context.Context, cid int64) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d", cid)
	cat := &Category{}
	err := c.DoGet(ctx, url, cat)
	return cat, err
}

func (c *Client) GetCategoryTranslated(ctx context.Context, cid int64, lang string) (*Category, error) {
	url := c.Endpoint("/solutions/categories/%d/%s", cid, lang)
	cat := &Category{}
	err := c.DoGet(ctx, url, cat)
	return cat, err
}

func (c *Client) ListCategories(ctx context.Context, lco *ListCategoriesOption) ([]*Category, bool, error) {
	url := c.Endpoint("/solutions/categories")
	categories := []*Category{}
	next, err := c.DoList(ctx, url, lco, &categories)
	return categories, next, err
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

func (c *Client) ListCategoriesTranslated(ctx context.Context, lang string, lco *ListCategoriesOption) ([]*Category, bool, error) {
	url := c.Domain + "/api/v2/solutions/categories/" + lang
	categories := []*Category{}
	next, err := c.DoList(ctx, url, lco, &categories)
	return categories, next, err
}

func (c *Client) IterCategoriesTranslated(ctx context.Context, lang string, lco *ListCategoriesOption, icf func(*Category) error) error {
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
		categories, next, err := c.ListCategoriesTranslated(ctx, lang, lco)
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

func (c *Client) CreateFolder(ctx context.Context, cid int64, folder *FolderCreate) (*Folder, error) {
	url := c.Endpoint("/solutions/categories/%d/folders", cid)
	result := &Folder{}
	if err := c.DoPost(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateFolderTranslated(ctx context.Context, fid int64, lang string, folder *FolderCreate) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d/%s", fid, lang)
	result := &Folder{}
	if err := c.DoPost(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateFolder(ctx context.Context, fid int64, folder *FolderUpdate) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d", fid)
	result := &Folder{}
	if err := c.DoPut(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateFolderTranslated(ctx context.Context, fid int64, lang string, folder *FolderUpdate) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d/%s", fid, lang)
	result := &Folder{}
	if err := c.DoPut(ctx, url, folder, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetFolder(ctx context.Context, fid int64) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d", fid)
	folder := &Folder{}
	err := c.DoGet(ctx, url, folder)
	return folder, err
}

func (c *Client) GetFolderTranslated(ctx context.Context, fid int64, lang string) (*Folder, error) {
	url := c.Endpoint("/solutions/folders/%d/%s", fid, lang)
	folder := &Folder{}
	err := c.DoGet(ctx, url, folder)
	return folder, err
}

func (c *Client) ListCategoryFolders(ctx context.Context, cid int64, lfo *ListFoldersOption) ([]*Folder, bool, error) {
	url := c.Endpoint("/solutions/categories/%d/folders", cid)
	folders := []*Folder{}
	next, err := c.DoList(ctx, url, lfo, &folders)
	return folders, next, err
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

func (c *Client) ListCategoryFoldersTranslated(ctx context.Context, cid int64, lang string, lfo *ListFoldersOption) ([]*Folder, bool, error) {
	url := c.Endpoint("/solutions/categories/%d/folders/%s", cid, lang)
	folders := []*Folder{}
	next, err := c.DoList(ctx, url, lfo, &folders)
	return folders, next, err
}

func (c *Client) IterCategoryFoldersTranslated(ctx context.Context, cid int64, lang string, lfo *ListFoldersOption, iff func(*Folder) error) error {
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
		folders, next, err := c.ListCategoryFoldersTranslated(ctx, cid, lang, lfo)
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

func (c *Client) ListSubFolders(ctx context.Context, fid int64, lfo *ListFoldersOption) ([]*Folder, bool, error) {
	url := c.Endpoint("/solutions/folders/%d/subfolders", fid)
	folders := []*Folder{}
	next, err := c.DoList(ctx, url, lfo, &folders)
	return folders, next, err
}

func (c *Client) IterSubFolders(ctx context.Context, fid int64, lfo *ListFoldersOption, iff func(*Folder) error) error {
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
		folders, next, err := c.ListSubFolders(ctx, fid, lfo)
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

func (c *Client) ListSubFoldersTranslated(ctx context.Context, fid int64, lang string, lfo *ListFoldersOption) ([]*Folder, bool, error) {
	url := c.Endpoint("/solutions/folders/%d/subfolders/%s", fid, lang)
	folders := []*Folder{}
	next, err := c.DoList(ctx, url, lfo, &folders)
	return folders, next, err
}

func (c *Client) IterSubFoldersTranslated(ctx context.Context, fid int64, lang string, lfo *ListFoldersOption, iff func(*Folder) error) error {
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
		folders, next, err := c.ListSubFoldersTranslated(ctx, fid, lang, lfo)
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

func (c *Client) CreateArticle(ctx context.Context, fid int64, article *ArticleCreate) (*Article, error) {
	url := c.Endpoint("/solutions/folders/%d/articles", fid)
	result := &Article{}
	if err := c.DoPost(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateArticleTranslated(ctx context.Context, aid int64, lang string, article *ArticleCreate) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d/%s", aid, lang)
	result := &Article{}
	if err := c.DoPost(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateArticle(ctx context.Context, aid int64, article *ArticleUpdate) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d", aid)
	result := &Article{}
	if err := c.DoPut(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateArticleTranslated(ctx context.Context, aid int64, lang string, article *ArticleUpdate) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d/%s", aid, lang)
	result := &Article{}
	if err := c.DoPut(ctx, url, article, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetArticle(ctx context.Context, aid int64) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d", aid)
	article := &Article{}
	err := c.DoGet(ctx, url, article)
	return article, err
}

func (c *Client) GetArticleTranslated(ctx context.Context, aid int64, lang string) (*Article, error) {
	url := c.Endpoint("/solutions/articles/%d/%s", aid, lang)
	article := &Article{}
	err := c.DoGet(ctx, url, article)
	return article, err
}

func (c *Client) ListFolderArticles(ctx context.Context, fid int64, lao *ListArticlesOption) ([]*Article, bool, error) {
	url := c.Endpoint("/solutions/folders/%d/articles", fid)
	articles := []*Article{}
	next, err := c.DoList(ctx, url, lao, &articles)
	return articles, next, err
}

func (c *Client) IterFolderArticles(ctx context.Context, fid int64, lao *ListArticlesOption, iaf func(*Article) error) error {
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

func (c *Client) ListFolderArticlesTranslated(ctx context.Context, fid int64, lang string, lao *ListArticlesOption) ([]*Article, bool, error) {
	url := c.Endpoint("/solutions/folders/%d/farticles/%s", fid, lang)
	articles := []*Article{}
	next, err := c.DoList(ctx, url, lao, &articles)
	return articles, next, err
}

func (c *Client) IterFolderArticlesTranslated(ctx context.Context, fid int64, lang string, lao *ListArticlesOption, iaf func(*Article) error) error {
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
		articles, next, err := c.ListFolderArticlesTranslated(ctx, fid, lang, lao)
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

func (c *Client) SearchArticles(ctx context.Context, keyword string) ([]*ArticleEx, error) {
	url := c.Endpoint("/search/solutions?term=%s", url.QueryEscape(keyword))
	articles := []*ArticleEx{}
	err := c.DoGet(ctx, url, &articles)
	return articles, err
}
