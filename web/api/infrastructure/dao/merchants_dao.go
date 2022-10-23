package dao

import (
	"log"
	"time"
)

type IMerchantDao interface {
	Save(merchant Merchants) error
	Get(address string) *Merchants
	GetAll() []Merchants
}

type merchantDao struct {
}

type Merchants struct {
	Id              uint      `db:"id"`
	Address         string    `db:"address"`
	ReceivedAddress string    `db:"received_address"`
	Deposit         uint      `db:"deposit"`
	CancelableTime  uint      `db:"cancelable_time"`
	Name            string    `db:"name"`
	Phonenumber     string    `db:"phonenumber"`
	MerchantAddress string    `db:"merchant_address"`
	CreateTs        time.Time `db:"create_ts"`
	UpdateTs        time.Time `db:"update_ts"`
	//DeleteTs        sql.NullTime `db:"delete_ts"`
}

func NewMerchantDao() IMerchantDao {
	return &merchantDao{}
}

func (o merchantDao) Save(merchant Merchants) error {
	dbmap := initDb()
	defer dbmap.Db.Close()
	merchant.CreateTs = time.Now()
	merchant.UpdateTs = time.Now()
	return dbmap.Insert(&merchant)
}

func (o merchantDao) GetAll() []Merchants {
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

func (o merchantDao) Get(address string) *Merchants {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var merchant Merchants
	err := dbmap.SelectOne(&merchant, "select * from merchants where address = ?", address)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &merchant
}
