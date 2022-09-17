package middleware

import (
	"stratosphaere-server/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			token := authHeader[len(BEARER_SCHEMA):]
			if token != "" {
				claims, err := util.ParseToken(token)
				if err == nil {
					c.Set("user", claims.UserID)
				}
			}
		}

		c.Next()
	}
}
