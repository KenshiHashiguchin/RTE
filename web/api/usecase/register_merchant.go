package usecase

import (
	"github.com/RTE/web/api/domain/repository"
)

type IRegisterMerchant interface {
	Register(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) error
}

type registerMerchant struct {
	repo repository.MerchantRepository
}

func NewRegisterMerchantsUseCase(repo repository.MerchantRepository) IRegisterMerchant {
	return &registerMerchant{repo}
}

func (t registerMerchant) Register(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) error {
	return t.repo.Register(address, receiveAddress, name, tel, merchantAddress, deposit, cancelableTime)
}
