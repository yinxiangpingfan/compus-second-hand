package utils

import (
	"compus-second-hand/global"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//关于鉴权

// 生成token
type JWTClaim struct {
	ID int
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	claims := &JWTClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.Configs.JWT.ExpireTime) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.Configs.JWT.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 用于鉴权的中间件
func ValidateToken(signedtoken string) (int, error) {
	token, err := jwt.ParseWithClaims(signedtoken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Configs.JWT.SecretKey), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return -1, errors.New("invalid token")
	}
	return claims.ID, nil
}
