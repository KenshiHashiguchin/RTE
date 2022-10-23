package repository

import "github.com/RTE/web/api/domain/model"

type MerchantRepository interface {
	CreateUser(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) model.Merchant
	GetMerchants() []model.Merchant
}
