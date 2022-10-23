package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IGetMerchant interface {
	GetMerchant(address string) *model.Merchant
}

type getMerchant struct {
	repo repository.MerchantRepository
}

func NewGetMerchantUseCase(repo repository.MerchantRepository) IGetMerchant {
	return &getMerchant{repo}
}

func (t getMerchant) GetMerchant(address string) *model.Merchant {
	return t.repo.GetMerchant(address)
}
