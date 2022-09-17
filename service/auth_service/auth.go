package auth_service

import (
	"github.com/go-playground/validator/v10"
)

type Auth struct {
	Username string `max=20,required`
	Password string `max=50,required`
}

func (a *Auth) Validate() error {
	validation := validator.New()
	return validation.Struct(a.Username)
}

func (a *Auth) Check() (bool, error) {

}
