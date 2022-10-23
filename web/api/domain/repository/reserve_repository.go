package repository

import "github.com/RTE/web/api/domain/model"

type ReserveRepository interface {
	Save(reservation model.Reservation) error
	Get(address string) []model.Reservation
	GetByMerchantAddress(address string) []model.Reservation
}
