package api

import (
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"
	"stratosphaere-server/service/auth_service"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	auth := auth_service.Auth{Username: username, Password: password}

	err := auth.Validate()
	if err != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
	}

	exists, err := auth.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
	}
	if !exists {
		appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_INVALID_CREDENTIALS, nil)
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_TOKEN_FAIL, nil)
	}

	appG.Response(http.StatusOK, exception.SUCCESS, map[string]string{
		"token": token,
	})
}
