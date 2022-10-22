package model

import (
	"github.com/RTE/web/api/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTExpHour time.Duration = 24

type UserClaims struct {
	address string `json:"address,omitempty"`
	jwt.StandardClaims
}

func NewUserClaims(address string) UserClaims {
	return UserClaims{
		address,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * JWTExpHour).Unix(),
		},
	}
}

func GetAuthUserByTokenString(tokenString string) (UserClaims, error) {
	auth := &UserClaims{}
	_, err := jwt.ParseWithClaims(tokenString, auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.Env("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return UserClaims{}, err
	}

	return *auth, nil
}
