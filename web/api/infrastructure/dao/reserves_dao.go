package dao

import (
	"log"
	"time"
)

type IReserveDao interface {
	Save(reserve Reserves) error
	GetAllByAddress(address string) []Reserves
}

type reserveDao struct {
}

type Reserves struct {
	Id              string    `db:"id"`
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

func (o reserveDao) GetAllByAddress(address string) []Reserves {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var reserves []Reserves
	_, err := dbmap.Select(&reserves, "select * from reserves where reserved_address = ?", address)
	if err != nil {
		log.Print(err)
		return nil
	}
	return reserves
}
