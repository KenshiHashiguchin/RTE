package dao

import (
	"log"
	"time"
)

type IMerchantDao interface {
	Save(merchant Merchant) error
	Get(address string) *Merchant
	GetById(id uint) (*Merchant, error)
	GetAll() []Merchant
}

type merchantDao struct {
}

type Merchant struct {
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
}

func NewMerchantDao() IMerchantDao {
	return &merchantDao{}
}

func (o merchantDao) Save(merchant Merchant) error {
	dbmap := initDb()
	defer dbmap.Db.Close()
	merchant.CreateTs = time.Now()
	merchant.UpdateTs = time.Now()
	return dbmap.Insert(&merchant)
}

func (o merchantDao) GetAll() []Merchant {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var merchants []Merchant
	_, err := dbmap.Select(&merchants, "select * from merchants")
	if err != nil {
		log.Print(err)
		return nil
	}
	return merchants
}

func (o merchantDao) Get(address string) *Merchant {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var merchant Merchant
	err := dbmap.SelectOne(&merchant, "select * from merchants where address = ?", address)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &merchant
}

func (o merchantDao) GetById(id uint) (*Merchant, error) {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var merchant Merchant
	err := dbmap.SelectOne(&merchant, "select * from merchants where id = ?", id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &merchant, nil
}
