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

func GenerateToken(user_id uint16, username string, expires bool) (string, error) {
	expireTime := time.Now().Add(24 * 30 * time.Hour)
	var registeredClaims jwt.RegisteredClaims
	if expires {
		registeredClaims = jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "solis",
		}
	} else {
		registeredClaims = jwt.RegisteredClaims{
			Issuer: "solis",
		}
	}

	authClaims := &AuthClaims{
		int(user_id),
		username,
		registeredClaims,
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
