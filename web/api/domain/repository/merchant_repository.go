package repository

import "github.com/RTE/web/api/domain/model"

type MerchantRepository interface {
	Register(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) error
	GetMerchants() []model.Merchant
}
