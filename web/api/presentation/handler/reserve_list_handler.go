package handler

import (
	"github.com/RTE/web/api/infrastructure/dao"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reserveListHandler struct {
	usecase usecase.IGetReservation
}

func newReserveListHandler() *reserveListHandler {
	rDao := dao.NewReserveDao()
	mDao := dao.NewMerchantDao()
	repo := repositoryImpl.NewReserveRepository(rDao, mDao)
	return &reserveListHandler{
		usecase: usecase.NewGetReservationUseCase(repo),
	}
}

type resReserve struct {
	PaymentId   string      `json:"payment_id"`
	MerchantId  uint        `json:"merchant_id"`
	Surname     string      `json:"surname"`
	Firstname   string      `json:"first_name"`
	Phonenumber string      `json:"phonenumber"`
	Number      uint        `json:"number"`
	Merchant    resMerchant `json:"merchant"`
}

type resMerchant struct {
	ReceivedAddress string `json:"received_address"`
	Name            string `json:"name"`
	Tel             string `json:"phonenumber"`
	MerchantAddress string `json:"merchant_address"`
	Deposit         uint   `json:"deposit"`
	CancelableTime  uint   `json:"cancelable_time"`
}

func HandleReserveList(c *gin.Context) {
	user, err := AuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	reservations := newReserveListHandler().usecase.GetReservation(user.GetAddress())
	var res []resReserve
	for _, r := range reservations {
		merchant := resMerchant{
			r.GetMerchant().GetReceivedAddress(),
			r.GetMerchant().GetName(),
			r.GetMerchant().GetTel(),
			r.GetMerchant().GetMerchantAddress(),
			r.GetMerchant().GetDeposit(),
			r.GetMerchant().GetCancelableTime(),
		}

		res = append(res, resReserve{
			r.GetPaymentId(),
			r.GetMerchantId(),
			r.GetSurname(),
			r.GetFirstname(),
			r.GetTel(),
			r.GetNumber(),
			merchant,
		})
	}
	c.JSON(http.StatusOK, gin.H{"reservations": res})
}
