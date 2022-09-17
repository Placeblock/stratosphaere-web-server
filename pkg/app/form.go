package app

import (
	"net/http"
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	if c.Bind(form) != nil {
		return http.StatusBadRequest, exception.INVALID_PARAMS
	}

	valid := validator.Validate{}
	if valid.Struct(form) != nil {
		return http.StatusBadRequest, exception.INVALID_PARAMS
	}

	return http.StatusOK, exception.SUCCESS
}
