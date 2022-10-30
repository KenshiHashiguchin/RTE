package model

type Merchant struct {
	id              uint
	address         string
	receivedAddress string
	name            string
	merchantAddress string
	tel             string
	deposit         uint
	cancelableTime  uint
}

func NewMerchant(id uint, address, receivedAddress, name, merchantAddress, tel string, deposit, cancelableTime uint) Merchant {
	return Merchant{
		id:              id,
		address:         address,
		receivedAddress: receivedAddress,
		name:            name,
		merchantAddress: merchantAddress,
		tel:             tel,
		deposit:         deposit,
		cancelableTime:  cancelableTime,
	}
}

func (m Merchant) GetId() uint {
	return m.id
}

func (m Merchant) GetAddress() string {
	return m.address
}

func (m Merchant) GetReceivedAddress() string {
	return m.receivedAddress
}

func (m Merchant) GetMerchantAddress() string {
	return m.merchantAddress
}

func (m Merchant) GetName() string {
	return m.name
}

func (m Merchant) GetTel() string {
	return m.tel
}

func (m Merchant) GetDeposit() uint {
	return m.deposit
}

func (m Merchant) GetCancelableTime() uint {
	return m.cancelableTime
}
