package v1

import (
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"
	"stratosphaere-server/service/auth_service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type auth struct {
	Username string `validate:"required,max=50"`
	Password string `validate:"required,max=50"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	auth := auth{Username: username, Password: password}

	validation := validator.New()
	valid := validation.Struct(auth)
	if valid != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	exists, id, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_INVALID_CREDENTIALS, nil)
		return
	}

	token, err := util.GenerateToken(id, username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_TOKEN_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, map[string]string{
		"token": token,
	})
}
