package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID            *uint16        `json:"id,omitempty"`
	Title         *string        `json:"title,omitempty" gorm:"default:'Unbekannter Titel'"`
	Description   *string        `json:"description,omitempty" gorm:"default:'Unbekannte Beschreibung'"`
	CoverImageUrl *string        `json:"cover_image_url,omitempty" gorm:"default:'https://cdn.pixabay.com/photo/2017/06/17/10/55/hot-air-balloon-2411851_960_720.jpg'"`
	Author        *string        `json:"author,omitempty"`
	Published     *bool          `json:"published,omitempty" gorm:"default:0"`
	PublishDate   *time.Time     `json:"publish_date" gorm:"default:0001-01-01 00:00:00 +0000 UTC"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Content       *string        `json:"content,omitempty"`
}

func (d Article) MarshalJSON() ([]byte, error) {
	type Alias Article
	return json.Marshal(&struct {
		*Alias
		PublishDate string  `json:"publish_date,omitempty"`
		UpdatedAt   string  `json:"updated_at,omitempty"`
		DeletedAt   *string `json:"deleted_at,omitempty"`
	}{
		Alias:       (*Alias)(&d),
		PublishDate: (*d.PublishDate).UTC().Format(time.RFC3339),
		UpdatedAt:   (*d.UpdatedAt).UTC().Format(time.RFC3339),
		DeletedAt:   nil,
	})
}

func (a Article) GetIDChunk(chunkParameters *GetArticlesParams) ([]*uint16, error) {
	var articleids []*uint16
	err := db.Model(&Article{}).
		Select("articles.id").
		//pub & (state=true) | unpub & (state=false)
		Where("(articles.published AND ?) OR (NOT articles.published AND ?)",
			chunkParameters.Published, chunkParameters.Unpublished).
		Order("articles.published desc").Order("articles.publish_date desc").Order("articles.updated_at desc").
		Offset(*chunkParameters.Offset).Limit(*chunkParameters.Limit).
		Joins("JOIN users ON users.username = articles.author").
		Find(&articleids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleids, nil
}

func GetAllLastModified() (time.Time, error) {
	var updatedAt time.Time
	var subquery1 = db.Model(&Article{}).Select("articles.updated_at AS modified_date")
	var subquery2 = db.Model(&Article{}).Unscoped().Select("articles.deleted_at AS modified_date")
	var model = db.Table("(?) as m", db.Model(&Article{}).Raw("? UNION ?", subquery1, subquery2))
	err := model.Order("modified_date DESC").Limit(1).Find(&updatedAt).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return time.Now(), err
	}
	return updatedAt.UTC(), nil
}

func (a Article) Get(fields []string) (*Article, error) {
	var article Article
	for i, field := range fields {
		fields[i] = "articles." + field
	}
	err := db.Select(fields).Where("articles.id = ?", a.ID).Joins("JOIN users ON users.username = articles.author").First(&article).Error
	if err != nil {
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

func (a Article) Visibility(visible bool) (time.Time, error) {
	var updateInterface = map[string]interface{}{"published": visible}
	var publishTime time.Time
	if visible {
		publishTime = time.Now()
	} else {
		publishTime = time.UnixMilli(0)
	}
	updateInterface["publish_date"] = publishTime
	if err := db.Model(&Article{}).Where("id = ?", a.ID).Updates(updateInterface).Error; err != nil {
		return time.UnixMilli(0), err
	}
	return publishTime, nil
}

func (a *Article) Create() error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	if err := db.Where("id = ?", a.ID).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
