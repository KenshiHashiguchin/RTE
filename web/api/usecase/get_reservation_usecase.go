package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IGetReservation interface {
	GetReservation(address string) []model.Reservation
}

type getReservation struct {
	repo repository.ReserveRepository
}

func NewGetReservationUseCase(repo repository.ReserveRepository) IGetReservation {
	return &getReservation{repo}
}

func (t getReservation) GetReservation(address string) []model.Reservation {
	return t.repo.Get(address)
}
