package models

import "github.com/jinzhu/gorm"

type Article struct {
	ID            uint16 `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Author        string `json:"author"`
	Published     bool   `json:"published"`
}

func ExistArticleByID(id uint16) (bool, error) {
	var exists bool
	err := db.Select("count(*) > 0").Where("id = ?", id).First(&exists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func GetArticles(page uint16, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Where(maps).Offset(page).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id uint16) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
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

func AddArticle(article Article) error {
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

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
