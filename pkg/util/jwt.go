package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

type AuthClaims struct {
	UserID int `json:"userid"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id uint16) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)

	authClaims := &AuthClaims{
		int(user_id),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "solis",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	signedString, err := tokenClaims.SignedString(jwtSecret)

	fmt.Println("Generated Token: " + signedString)

	testparse, err := ParseToken(signedString)
	fmt.Println("TestParse: " + fmt.Sprint(err))
	fmt.Println(testparse)

	return signedString, err
}

func ParseToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	fmt.Println(token.Valid)
	fmt.Println(token.Claims)
	if token != nil {
		if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
