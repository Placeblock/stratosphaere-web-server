package article_serivce

import (
	"stratosphaere-server/models"
)

type Article struct {
	ID            uint16
	Title         string
	Description   string
	Content       string
	CoverImageUrl string
	Published     bool
	Author        uint16
	AuthorName    string
}

func (a *Article) Add() (uint16, error) {
	article := models.Article{
		ID:            a.ID,
		Title:         a.Title,
		Description:   a.Description,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		Author:        a.Author,
		Published:     a.Published,
	}
	err := models.AddArticle(&article)
	if err != nil {
		return 0, err
	}
	return article.ID, nil
}

func (a *Article) Edit() error {
	return models.EditArticle(models.Article{
		ID:            a.ID,
		Title:         a.Title,
		Description:   a.Description,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		Author:        a.Author,
		Published:     a.Published,
	})
}

func (a *Article) Get() (*models.Article, error) {
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *Article) GetAll(pageNum, pageSize int) ([]*models.Article, error) {
	articles, err := models.GetArticles(pageNum, pageSize, a)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) Exists() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) getMaps() map[string]interface{} {
	return make(map[string]interface{})
}

func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal(a.getMaps())
}
