package models

import (
	"errors"

	"gorm.io/gorm"
)

type Article struct {
	ID            uint16 `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Author        uint16 `json:"author"`
	Published     bool   `json:"published"`
}

var article_cache map[string]Article = make(map[string]Article)

func ExistArticleByID(id uint16) (bool, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func GetArticles(page int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Where(maps).Offset(page * pageSize).Limit(pageSize).Joins("JOIN users ON users.id = articles.author").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id uint16) (*Article, error) {
	var article Article
	err := db.Where("articles.id = ?", id).Joins("JOIN users ON users.id = articles.author").First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(article Article) error {
	if err := db.Model(&Article{}).Where("id = ?", article.ID).Updates(article).Error; err != nil {
		return err
	}

	return nil
}

func AddArticle(article *Article) error {
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id uint16) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
