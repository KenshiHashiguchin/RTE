package kvs

import (
	"time"
)

type IMerchantToken interface {
	SetToken(address string, value string)
	GetToken(address string) string
}

type merchantToken struct {
}

func NewMerchantToken() IMerchantToken {
	return &merchantToken{}
}

const merchantTokenKey = "merchant_token:"
const merchantTokenExpire = 86400

func (o merchantToken) SetToken(address string, value string) {
	err := conn.Set(ctx, merchantTokenKey+address, value, merchantTokenExpire*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func (o merchantToken) GetToken(address string) string {
	val, _ := conn.Get(ctx, merchantTokenKey+address).Result()
	return val
}
