package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

type AuthClaims struct {
	UserID   int    `json:"userid"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id uint16, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)

	authClaims := &AuthClaims{
		int(user_id),
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "solis",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	signedString, err := tokenClaims.SignedString(jwtSecret)
	return signedString, err
}

func ParseToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
