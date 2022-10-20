package model

import (
	"crypto/rand"
	"errors"
)

const tokenLength = 24

type UserToken struct {
	address string
	token   string
}

func (o UserToken) GetAddress() string {
	return o.address
}
func (o UserToken) GetToken() string {
	return o.token
}

func NewUserToken(address string) *UserToken {
	t, _ := generateUserToken()
	return &UserToken{address: address, token: t}
}

func generateUserToken() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, tokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
