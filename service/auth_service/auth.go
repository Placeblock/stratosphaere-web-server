package auth_service

import (
	"stratosphaere-server/models"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, uint16, error) {
	valid, id, err := models.CheckAuth(a.Username, a.Password)
	return valid, id, err
}
