package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IGetMerchants interface {
	GetMerchants() []model.Merchant
}

type getMerchants struct {
	repo repository.MerchantRepository
}

func NewGetMerchantsUseCase(repo repository.MerchantRepository) IGetMerchants {
	return &getMerchants{repo}
}

func (t getMerchants) GetMerchants() []model.Merchant {
	return t.repo.GetMerchants()
}
