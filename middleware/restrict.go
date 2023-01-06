package middleware

import (
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
)

func RestrictBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		username, exists := c.Get("user")
		if !exists || username == "" {
			appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func RestrictWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		username, exists := c.Get("user")
		if !exists || username != "webhook" {
			appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
