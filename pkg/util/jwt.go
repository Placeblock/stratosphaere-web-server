package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

type Claims struct {
	UserID   uint16 `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id uint16, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)

	fmt.Println("Generating Token...")
	fmt.Println(user_id)
	fmt.Println(username)
	fmt.Println(password)

	claims := Claims{
		user_id,
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "stratosphaere",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
