package repositoryImpl

import (
	"github.com/RTE/web/api/domain/model"
	repo "github.com/RTE/web/api/domain/repository"
	"github.com/RTE/web/api/infrastructure/dao"
)

type merchantImpl struct {
	dao dao.IMerchantDao
}

func NewMerchantRepository(dao dao.IMerchantDao) repo.MerchantRepository {
	return &merchantImpl{dao: dao}
}

func (u merchantImpl) CreateUser(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) model.Merchant {

	return model.Merchant{}
}

func (u merchantImpl) GetMerchant(address string) *model.Merchant {
	data := u.dao.Get(address)
	if data == nil {
		return nil
	}
	merchant := model.NewMerchant(
		data.Address,
		data.ReceivedAddress,
		data.Name,
		data.MerchantAddress,
		data.Phonenumber,
		data.Deposit,
		data.CancelableTime,
	)

	return &merchant
}

func (u merchantImpl) GetMerchants() []model.Merchant {
	merchantData := u.dao.GetAll()
	var merchants []model.Merchant
	for _, data := range merchantData {
		merchant := model.NewMerchant(
			data.Address,
			data.ReceivedAddress,
			data.Name,
			data.MerchantAddress,
			data.Phonenumber,
			data.Deposit,
			data.CancelableTime,
		)
		merchants = append(merchants, merchant)
	}

	return merchants
}

func (u merchantImpl) Register(address, receiveAddress, name, tel, merchantAddress string, deposit, cancelableTime uint) error {
	merchant := dao.Merchant{
		Address:         address,
		ReceivedAddress: receiveAddress,
		Deposit:         deposit,
		CancelableTime:  cancelableTime,
		Name:            name,
		Phonenumber:     tel,
		MerchantAddress: merchantAddress,
	}
	return u.dao.Save(merchant)
}
