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
	merchantData := r.merchantDao.GetAll()

	var reservations []model.Reservation
	for _, data := range reservationData {
		for _, merchantD := range merchantData {
			if merchantD.Id == data.MerchantId {
				merchant := model.NewMerchant(
					merchantD.Id,
					merchantD.Address,
					merchantD.ReceivedAddress,
					merchantD.Name,
					merchantD.MerchantAddress,
					merchantD.Phonenumber,
					merchantD.Deposit,
					merchantD.CancelableTime,
				)
				reservations = append(reservations, model.NewReservation(
					data.Id,
					data.MerchantId,
					data.ReservedAddress,
					data.Surname,
					data.Firstname,
					data.Phonenumber,
					data.Number,
					merchant,
				))
			}
		}
	}
	return reservations
}

func (r reserveImpl) GetByMerchantAddress(address string) []model.Reservation {
	merchantData := r.merchantDao.Get(address)
	if merchantData == nil {
		return nil
	}

	merchant := model.NewMerchant(
		merchantData.Id,
		merchantData.Address,
		merchantData.ReceivedAddress,
		merchantData.Name,
		merchantData.MerchantAddress,
		merchantData.Phonenumber,
		merchantData.Deposit,
		merchantData.CancelableTime,
	)

	reservationData := r.reserveDao.GetAllByMerchantId(merchantData.Id)
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
			merchant,
		))
	}
	return reservations
}
