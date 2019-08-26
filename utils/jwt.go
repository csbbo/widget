package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = "lwaids"
var Issuer = "hasani"
type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid int `json:"uid"`
	Username string `json:"username"`
}

func CreateJWT(uid int, username string) (string,error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    Issuer,
		},
		uid,
		username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(SecretKey))
	return jwt,err
}

func ParseJWT(tokenSrt string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}