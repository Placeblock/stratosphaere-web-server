package util

import (
	"stratosphaere-server/pkg/setting"

	"github.com/go-playground/validator"
)

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	Validate = validator.New()
}
