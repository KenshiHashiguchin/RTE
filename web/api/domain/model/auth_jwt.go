package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTExpHour time.Duration = 24

type UserClaims struct {
	Address string `json:"address,omitempty"`
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

func (u UserClaims) GetAddress() string {
	return u.Address
}

func GetAuthUserByTokenString(tokenString, key string) (UserClaims, error) {
	auth := &UserClaims{}
	_, err := jwt.ParseWithClaims(tokenString, auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return UserClaims{}, err
	}

	return *auth, nil
}
