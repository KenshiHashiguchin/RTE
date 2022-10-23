package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IMerchantReservation interface {
	GetReservation(address string) []model.Reservation
}

type merchantReservation struct {
	repo repository.ReserveRepository
}

func NewMerchantReservationUseCase(repo repository.ReserveRepository) IMerchantReservation {
	return &merchantReservation{repo}
}

func (t merchantReservation) GetReservation(address string) []model.Reservation {
	return t.repo.GetByMerchantAddress(address)
}
