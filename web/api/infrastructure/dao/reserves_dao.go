package dao

import (
	"time"
)

type IReserveDao interface {
	Save(reserve Reserves) error
}

type reserveDao struct {
}

type Reserves struct {
	Id              uint      `db:"id"`
	ReservedAddress string    `db:"reserved_address"`
	MerchantId      uint      `db:"merchant_id"`
	Surname         string    `db:"surname"`
	Firstname       string    `db:"firstname"`
	Phonenumber     string    `db:"phonenumber"`
	Number          uint      `db:"number"`
	CreateTs        time.Time `db:"create_ts"`
	UpdateTs        time.Time `db:"update_ts"`
}

func NewReserveDao() IReserveDao {
	return &reserveDao{}
}

func (o reserveDao) Save(reserve Reserves) error {
	dbmap := initDb()
	defer dbmap.Db.Close()
	reserve.CreateTs = time.Now()
	reserve.UpdateTs = time.Now()
	return dbmap.Insert(&reserve)
}
