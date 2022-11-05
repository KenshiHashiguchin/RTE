package usecase

import (
	"errors"
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IReserve interface {
	Reserve(merchantAddress, reservedAddress, surname, firstname, tel string, number uint) error
}

type reserve struct {
	reserveRepo  repository.ReserveRepository
	merchantRepo repository.MerchantRepository
}

func NewReserveUseCase(reserveRepo repository.ReserveRepository, merchantRepo repository.MerchantRepository) IReserve {
	return &reserve{reserveRepo, merchantRepo}
}

func (t reserve) Reserve(merchantAddress, reservedAddress, surname, firstname, tel string, number uint) error {
	paymentId, _ := model.GeneratePaymentId()
	merchant := t.merchantRepo.GetMerchant(merchantAddress)
	if merchant == nil {
		return errors.New("Don't exist")
	}
	reservation := model.NewReservation(paymentId, merchant.GetId(), reservedAddress, surname, firstname, tel, number, *merchant)
	return t.reserveRepo.Save(reservation)
}
