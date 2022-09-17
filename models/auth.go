package models

import (
	"stratosphaere-server/pkg/util"

	"gorm.io/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key,autoIncrement" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "users"
}

func CheckAuth(username, password string) (bool, uint16, error) {
	var auth Auth
	err := db.Select("id, password").Where(Auth{Username: username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, 0, err
	}

	return util.CompareHash(auth.Password, []byte(password)), uint16(auth.ID), nil
}
