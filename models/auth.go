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

var user_cache map[uint16]Auth = make(map[uint16]Auth)

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

func GetName(userid uint16) (string, error) {
	if user, ok := user_cache[userid]; ok {
		return user.Username, nil
	}
	var user Auth
	err := db.Where(Auth{ID: userid}).First(&user).Error
	if err != nil {
		return "", err
	}
	user_cache[userid] = user
	return user.Username, nil
}
