package middleware

import (
	"net/http"
	"stratosphaere-server/pkg/exception"
	"stratosphaere-server/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = exception.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = exception.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != exception.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  exception.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
