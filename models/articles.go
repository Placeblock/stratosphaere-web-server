package models

import (
	"errors"

	"gorm.io/gorm"
)

type Article struct {
	ID            uint16 `json:"id"`
	Title         string `json:"title" gorm:"default:'Unbekannter Titel'"`
	Description   string `json:"description" gorm:"default:'Unbekannte Beschreibung'"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url" gorm:"default:'https://cdn.pixabay.com/photo/2017/06/17/10/55/hot-air-balloon-2411851_960_720.jpg'"`
	Author        string `json:"author"`
	Published     bool   `json:"published" gorm:"default:0"`
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

func (a Article) GetAll(offset int, amount int) ([]*Article, error) {
	var articles []*Article
	err := db.Where(a).Offset(offset).Limit(amount).Joins("JOIN users ON users.username = articles.author").Find(&articles).Error
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
