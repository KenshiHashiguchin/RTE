package model

import (
	"crypto/rand"
	"errors"
	"time"
)

type Reservation struct {
	paymentId       string
	merchantId      uint
	reservedAddress string
	surname         string
	firstname       string
	tel             string
	number          uint
	reserveTs       time.Time
	merchant        Merchant
}

func NewReservation(paymentId string, merchantId uint, reservedAddress, surname, firstname, tel string, number uint, reserveTs time.Time, merchant Merchant) Reservation {
	return Reservation{
		paymentId,
		merchantId,
		reservedAddress,
		surname,
		firstname,
		tel,
		number,
		reserveTs,
		merchant,
	}
}

func (r Reservation) GetPaymentId() string {
	return r.paymentId
}

func (r Reservation) GetMerchantId() uint {
	return r.merchantId
}

func (r Reservation) GetReservedAddress() string {
	return r.reservedAddress
}

func (r Reservation) GetSurname() string {
	return r.surname
}

func (r Reservation) GetFirstname() string {
	return r.firstname
}

func (r Reservation) GetTel() string {
	return r.tel
}

func (r Reservation) GetNumber() uint {
	return r.number
}

func (r Reservation) GetReserveTs() time.Time {
	return r.reserveTs
}

func (r Reservation) GetMerchant() Merchant {
	return r.merchant
}

func GeneratePaymentId() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 12)
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
