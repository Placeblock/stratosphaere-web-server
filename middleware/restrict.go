package middleware

import (
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
)

func Restrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		userid, exists := c.Get("user")
		if !exists || userid == "" {
			appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
