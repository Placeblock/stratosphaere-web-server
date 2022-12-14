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
	Username string `json:"username" binding:"required" validate:"required,max=50"`
	Password string `json:"password" binding:"required" validate:"required,max=50"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	auth := auth{}
	if c.BindJSON(&auth) != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	validation := validator.New()
	valid := validation.Struct(auth)
	if valid != nil {
		appG.Response(http.StatusBadRequest, exception.INVALID_PARAMS, nil)
		return
	}

	authModel := models.Auth{Username: auth.Username, Password: auth.Password}
	exists, err := authModel.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusBadRequest, exception.ERROR_AUTH_INVALID_CREDENTIALS, nil)
		return
	}

	token, err := util.GenerateToken(authModel.ID, authModel.Username, true)
	if err != nil {
		appG.Response(http.StatusInternalServerError, exception.ERROR_AUTH_TOKEN_FAIL, nil)
		return
	}
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("authToken", token, 60*60*24*30, "/", "", true, false)
	appG.Response(http.StatusOK, exception.SUCCESS, nil)
}
