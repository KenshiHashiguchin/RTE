package usecase

import (
	"errors"
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
	"log"
	"time"
)

type IReserve interface {
	Reserve(paymentId, merchantAddress, reservedAddress, surname, firstname, tel string, number uint, reserveTs time.Time) error
}

type reserve struct {
	reserveRepo  repository.ReserveRepository
	merchantRepo repository.MerchantRepository
}

func NewReserveUseCase(reserveRepo repository.ReserveRepository, merchantRepo repository.MerchantRepository) IReserve {
	return &reserve{reserveRepo, merchantRepo}
}

func (t reserve) Reserve(paymentId, merchantAddress, reservedAddress, surname, firstname, tel string, number uint, reserveTs time.Time) error {
	merchant := t.merchantRepo.GetMerchant(merchantAddress)
	if merchant == nil {
		log.Print("merchant don't exist")
		return errors.New("Don't exist")
	}
	reservation := model.NewReservation(paymentId, merchant.GetId(), reservedAddress, surname, firstname, tel, number, reserveTs, *merchant)
	return t.reserveRepo.Save(reservation)
}
