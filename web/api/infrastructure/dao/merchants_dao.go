package dao

import (
	"database/sql"
	"log"
	"time"
)

type IMerchantDao interface {
	Save()
	GetAll() []Merchants
}

type merchantDao struct {
}

type Merchants struct {
	Id              uint         `db:"id"`
	Address         string       `db:"address"`
	ReceivedAddress string       `db:"received_address"`
	Deposit         uint         `db:"deposit"`
	CancelableTime  uint         `db:"cancelable_time"`
	Name            string       `db:"name"`
	Phonenumber     string       `db:"phonenumber"`
	MerchantAddress string       `db:"merchant_address"`
	CreateTs        time.Time    `db:"create_ts"`
	UpdateTs        time.Time    `db:"update_ts"`
	DeleteTs        sql.NullTime `db:"delete_ts"`
}

func NewMerchantDao() IMerchantDao {
	return &merchantDao{}
}

func (o merchantDao) Save() {
	dbmap := initDb()
	defer dbmap.Db.Close()
}

func (o merchantDao) GetAll() []Merchants {
	log.Print("GetAll")

	dbmap := initDb()
	defer dbmap.Db.Close()

	var merchants []Merchants

	_, err := dbmap.Select(&merchants, "select * from merchants")
	if err != nil {
		log.Print(err)
		return nil
	}
	return merchants
}
