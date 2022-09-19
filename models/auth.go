package models

import (
	"stratosphaere-server/pkg/util"

	"gorm.io/gorm"
)

type Auth struct {
	ID       uint16 `gorm:"primary_key,autoIncrement" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "users"
}

func (a Auth) Check() (bool, error) {
	var auth Auth
	err := db.Select("id, password").Where(Auth{Username: a.Username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	a.ID = uint16(auth.ID)
	return util.CompareHash(auth.Password, []byte(a.Password)), nil
}
