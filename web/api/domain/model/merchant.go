package model

type Merchant struct {
	address         string
	receivedAddress string
	name            string
	merchantAddress string
	tel             string
	deposit         uint
	cancelableTime  uint
}

func NewMerchant(address, receivedAddress, name, merchantAddress, tel string, deposit, cancelableTime uint) Merchant {
	return Merchant{
		address:         address,
		receivedAddress: receivedAddress,
		name:            name,
		merchantAddress: merchantAddress,
		tel:             tel,
		deposit:         deposit,
		cancelableTime:  cancelableTime,
	}
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
