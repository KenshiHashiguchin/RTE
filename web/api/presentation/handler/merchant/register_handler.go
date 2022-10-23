package merchant

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/RTE/web/api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IRegisterHandler interface {
	HandleReserve(c *gin.Context)
}

type registerHandler struct {
	usecase usecase.IAuth
}

func newRegisterHandler() *registerHandler {
	dao := kvs.NewUserToken()
	return &registerHandler{
		usecase: usecase.NewAuthUseCase(dao),
	}
}

func HandleRegister(c *gin.Context) {
	//user, err := handler.AuthUser(c)
	_, err := handler.AuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}

	// request
	req := request.Register{
		ReceivedAddress: c.PostForm("received_address"),
		Name:            c.PostForm("name"),
		Tel:             c.PostForm("tel"),
		MerchantAddress: c.PostForm("merchant_address"),
		Deposit:         util.ConvertStringToUint(c.PostForm("deposit")),
		CancelableTime:  util.ConvertStringToUint(c.PostForm("cancelable_time")),
	}

	ok, errMessage := req.Validate()
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errMessage})
		return
	}

	// save merchants

	c.JSON(http.StatusOK, gin.H{})
}
