package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type LiveData struct {
	Latitude    float32    `json:"latitude"`
	Longitude   float32    `json:"longitude"`
	Altitude    int32      `json:"altitude"`
	Temperature float32    `json:"temperature"`
	Pressure    int16      `json:"pressure"`
	Humidity    int16      `json:"humidity"`
	CreatedAt   *time.Time `json:"time"`
}

func (d LiveData) MarshalJSON() ([]byte, error) {
	type Alias LiveData
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"time"`
	}{
		Alias:     (*Alias)(&d),
		CreatedAt: (*d.CreatedAt).UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}

func (d *LiveData) Create() error {
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func GetLiveData(since time.Time) ([]LiveData, error) {
	var liveData []LiveData
	err := db.Model(&LiveData{}).
		Where("created_at > ?", since.UTC().Format("2006-01-02 15:04:05.000")).
		Order("created_at asc").
		Find(&liveData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return liveData, nil
}
