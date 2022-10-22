package kvs

import (
	"context"
	"time"
)

type IUserToken interface {
	SetToken(address string, value string)
	GetToken(address string) string
}

type token struct {
}

func NewUserToken() IUserToken {
	return &token{}
}

const tokenKey = "token:"
const tokenExpire = 86400

var ct = context.Background()

func (o token) SetToken(address string, value string) {
	err := conn.Set(ctx, tokenKey+address, value, tokenExpire*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func (o token) GetToken(address string) string {
	val, _ := conn.Get(ctx, tokenKey+address).Result()
	return val
}
