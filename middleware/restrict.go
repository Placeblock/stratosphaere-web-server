package middleware

import (
	"fmt"
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"

	"github.com/gin-gonic/gin"
)

func Restrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		userid, exists := c.Get("user")
		fmt.Println(exists)
		fmt.Println("USERID: " + fmt.Sprint(userid))
		fmt.Printf("Authenticated: %t\n", exists && userid != "")
		if !exists || userid == "" {
			appG.Response(http.StatusUnauthorized, exception.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
