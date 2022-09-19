package v1

import (
	"net/http"
	"stratosphaere-server/models"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"

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

	authModel := models.Auth{Username: username, Password: password}
	exists, err := authModel.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_INVALID_CREDENTIALS, nil)
		return
	}

	token, err := util.GenerateToken(authModel.ID, authModel.Username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_TOKEN_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, exception.SUCCESS, map[string]string{
		"token": token,
	})
}
