package repositoryImpl

import (
	"github.com/RTE/web/api/domain/model"
	repo "github.com/RTE/web/api/domain/repository"
	"github.com/RTE/web/api/infrastructure/dao"
)

type reserveImpl struct {
	reserveDao  dao.IReserveDao
	merchantDao dao.IMerchantDao
}

func NewReserveRepository(reserveDao dao.IReserveDao, merchantDao dao.IMerchantDao) repo.ReserveRepository {
	return &reserveImpl{reserveDao, merchantDao}
}

func (r reserveImpl) Save(reservation model.Reservation) error {
	_, err := r.merchantDao.GetById(reservation.GetMerchantId())
	if err != nil {
		return err
	}

	reserve := dao.Reserves{
		ReservedAddress: reservation.GetReservedAddress(),
		MerchantId:      reservation.GetMerchantId(),
		Surname:         reservation.GetSurname(),
		Firstname:       reservation.GetFirstname(),
		Number:          reservation.GetNumber(),
	}

	return r.reserveDao.Save(reserve)
}

func (r reserveImpl) Get(address string) []model.Reservation {
	reservationData := r.reserveDao.GetAllByAddress(address)
	var reservations []model.Reservation
	for _, data := range reservationData {
		reservations = append(reservations, model.NewReservation(
			data.Id,
			data.MerchantId,
			data.ReservedAddress,
			data.Surname,
			data.Firstname,
			data.Phonenumber,
			data.Number,
		))
	}
	return reservations
}

func (r reserveImpl) GetByMerchantAddress(address string) []model.Reservation {
	merchant := r.merchantDao.Get(address)
	if merchant == nil {
		return nil
	}

	reservationData := r.reserveDao.GetAllByMerchantId(merchant.Id)
	var reservations []model.Reservation
	for _, data := range reservationData {
		reservations = append(reservations, model.NewReservation(
			data.Id,
			data.MerchantId,
			data.ReservedAddress,
			data.Surname,
			data.Firstname,
			data.Phonenumber,
			data.Number,
		))
	}
	return reservations
}
