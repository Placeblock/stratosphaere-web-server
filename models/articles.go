package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ArticleMetadata struct {
	Title         string `json:"title" gorm:"default:'Unbekannter Titel'"`
	Description   string `json:"description" gorm:"default:'Unbekannte Beschreibung'"`
	CoverImageUrl string `json:"cover_image_url" gorm:"default:'https://cdn.pixabay.com/photo/2017/06/17/10/55/hot-air-balloon-2411851_960_720.jpg'"`
	Author        string `json:"author"`
	Published     bool   `json:"published" gorm:"default:0"`
	PublishDate   int    `json:"publish_date"`
	UpdatedAt     int
}

type Article struct {
	ID              uint16 `json:"id"`
	ArticleMetadata `gorm:"embedded" json:"metadata"`
	Content         string `json:"content"`
}

var article_cache map[string]Article = make(map[string]Article)

func (a Article) Exists() (bool, error) {
	var article Article
	err := db.Where("id = ?", a.ID).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

type GetAllResult struct {
	ID       uint16          `json:"id"`
	Metadata ArticleMetadata `gorm:"embedded" json:"metadata"`
}

func (a Article) GetAll(offset int, amount int, onlyPublished bool) ([]*GetAllResult, error) {
	var articles []*GetAllResult
	var model = db.Model(&Article{})
	if onlyPublished {
		model = model.Where("published = true")
	}
	err := model.Order("articles.published asc").Order("articles.publish_date desc").Offset(offset).Limit(amount).Joins("JOIN users ON users.username = articles.author").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func (a Article) Get() (*Article, error) {
	var article Article
	err := db.Where("articles.id = ?", a.ID).Joins("JOIN users ON users.username = articles.author").First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func (a Article) Edit() error {
	if err := db.Model(&Article{}).Where("id = ?", a.ID).Updates(a).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Visibility(visible bool) (int, error) {
	var updateInterface = map[string]interface{}{"published": visible}
	var updateTime = int(time.Now().Unix())
	if visible {
		updateInterface["publish_date"] = updateTime
	}
	if err := db.Model(&Article{}).Where("id = ?", a.ID).Updates(updateInterface).Error; err != nil {
		return 0, err
	}
	return updateTime, nil
}

func (a *Article) Add() error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete() error {
	if err := db.Where("id = ?", a.ID).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Count() (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Where(a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
