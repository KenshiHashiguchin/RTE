package handler

import (
	"github.com/RTE/web/api/infrastructure/dao"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type merchantHandler struct {
	usecase usecase.IGetMerchant
}

func newMerchantHandler() *merchantHandler {
	dao := dao.NewMerchantDao()
	repo := repositoryImpl.NewMerchantRepository(dao)
	return &merchantHandler{
		usecase: usecase.NewGetMerchantUseCase(repo),
	}
}

func HandleMerchant(c *gin.Context) {
	merchant := newMerchantHandler().usecase.GetMerchant(c.Param("address"))
	if merchant == nil {
		c.JSON(http.StatusOK, gin.H{"merchant": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"merchant": res{
		merchant.GetAddress(),
		merchant.GetReceivedAddress(),
		merchant.GetName(),
		merchant.GetMerchantAddress(),
		merchant.GetTel(),
		merchant.GetDeposit(),
		merchant.GetCancelableTime(),
	}})
}
