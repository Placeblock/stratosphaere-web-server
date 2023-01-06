package models

import (
	"time"

	"gorm.io/gorm"
)

type LiveData struct {
	latitude    float32   `json:"latitude"`
	longitude   float32   `json:"longitude"`
	altitude    int32     `json:"altitude"`
	temperature int16     `json:"temperature"`
	pressure    int16     `json:"pressure"`
	humidity    int16     `json:"humidity"`
	CreatedAt   time.Time `json:"time"`
}

func GetLiveData(since time.Time) ([]LiveData, error) {
	var liveData []LiveData
	err := db.Model(&LiveData{}).
		Where("created_at > ?", since.Format("2006-01-02 15:04:05")).
		Order("created_at asc").
		Find(&liveData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return liveData, nil
}
