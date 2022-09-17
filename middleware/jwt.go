package middleware

import (
	"fmt"
	"net/http"
	"stratosphaere-server/pkg/app"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		fmt.Println("JWT MIDDLEWARE")
		var code int

		code = exception.SUCCESS

		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		token := authHeader[len(BEARER_SCHEMA):]
		fmt.Printf("token %s", token)
		if token == "" {
			code = exception.ERROR_AUTH_TOKEN_MISSING
		} else {
			claims, err := util.ParseToken(token)
			fmt.Println(claims)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else {
				c.Set("user", claims.UserID)
			}
		}

		if code != exception.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, nil)

			c.Abort()
			return
		}

		c.Next()
	}
}
