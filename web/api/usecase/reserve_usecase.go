package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
)

type IReserve interface {
	Reserve(merchantId uint, reservedAddress, surname, firstname, tel string, number uint) error
}

type reserve struct {
	repo repository.ReserveRepository
}

func NewReserveUseCase(repo repository.ReserveRepository) IReserve {
	return &reserve{repo}
}

func (t reserve) Reserve(merchantId uint, reservedAddress, surname, firstname, tel string, number uint) error {
	reservation := model.NewReservation(merchantId, reservedAddress, surname, firstname, tel, number)
	return t.repo.Save(reservation)
}
