package repository

import "github.com/RTE/web/api/domain/model"

type ReserveRepository interface {
	Save(reservation model.Reservation) error
}
