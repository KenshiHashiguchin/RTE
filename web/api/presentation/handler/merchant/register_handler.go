package merchant

import (
	"github.com/RTE/web/api/infrastructure/dao"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerHandler struct {
	usecase usecase.IRegisterMerchant
}

func newRegisterHandler() *registerHandler {
	dao := dao.NewMerchantDao()
	repo := repositoryImpl.NewMerchantRepository(dao)
	return &registerHandler{
		usecase: usecase.IRegisterMerchant(repo),
	}
}

func HandleRegister(c *gin.Context) {
	user, err := handler.AuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}

	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, errMessage := req.Validate()
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errMessage})
		return
	}

	err = newRegisterHandler().usecase.Register(
		user.GetAddress(),
		req.ReceivedAddress,
		req.Name,
		req.Tel,
		req.MerchantAddress,
		req.Deposit,
		req.CancelableTime,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{})
}
