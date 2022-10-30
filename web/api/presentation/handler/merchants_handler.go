package handler

import (
	"github.com/RTE/web/api/infrastructure/dao"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getStoresHandler struct {
	usecase usecase.IGetMerchants
}

func newGetStoresHandler() *getStoresHandler {
	dao := dao.NewMerchantDao()
	repo := repositoryImpl.NewMerchantRepository(dao)
	return &getStoresHandler{
		usecase: usecase.NewGetMerchantsUseCase(repo),
	}
}

type res struct {
	Address         string `json:"address"`
	ReceivedAddress string `json:"received_address"`
	Name            string `json:"name"`
	MerchantAddress string `json:"merchant_address"`
	Tel             string `json:"tel"`
	Deposit         uint   `json:"deposit"`
	CancelableTime  uint   `json:"cancelable_time"`
}

func HandleGetStores(c *gin.Context) {
	merchants := newGetStoresHandler().usecase.GetMerchants()
	var response []res
	for _, merchant := range merchants {
		response = append(response, res{
			merchant.GetAddress(),
			merchant.GetReceivedAddress(),
			merchant.GetName(),
			merchant.GetMerchantAddress(),
			merchant.GetTel(),
			merchant.GetDeposit(),
			merchant.GetCancelableTime(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"merchants": response})
}
