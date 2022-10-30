package handler

import (
	"github.com/RTE/web/api/infrastructure/dao"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/RTE/web/api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reserveHandler struct {
	UseCase usecase.IReserve
}

func newReserveHandler() *reserveHandler {
	rDao := dao.NewReserveDao()
	mDao := dao.NewMerchantDao()
	reserveRepo := repositoryImpl.NewReserveRepository(rDao, mDao)
	merchantRepo := repositoryImpl.NewMerchantRepository(mDao)
	return &reserveHandler{
		UseCase: usecase.NewReserveUseCase(reserveRepo, merchantRepo),
	}
}

func HandleReserve(c *gin.Context) {
	user, err := AuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	num := util.ConvertStringToUint(c.PostForm("number"))
	date := util.ConvertStringToTime(c.PostForm("time"))
	req := request.Reserve{
		c.PostForm("merchant_address"),
		c.PostForm("surname"),
		c.PostForm("firstname"),
		c.PostForm("phonenumber"),
		num,
		date,
	}
	ok, errMessage := req.Validate()
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errMessage})
	}

	// reserve
	err = newReserveHandler().UseCase.Reserve(req.MerchantAddress, user.GetAddress(), req.Surname, req.Firstname, req.Phonenumber, req.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{})
}
