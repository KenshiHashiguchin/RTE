package merchant

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newUserToken() *handler.UserToken {
	dao := kvs.NewMerchantToken() //merchant
	return &handler.UserToken{
		UseCase: usecase.NewGenerateUserTokenUsecase(dao),
	}
}

func HandleGetToken(c *gin.Context) {
	addr := c.Param("address")
	req := request.Token{addr}
	ok, err := req.Validate()
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	u := newUserToken().UseCase.GenerateUserToken(addr)
	c.JSON(http.StatusOK, gin.H{"token": u.GetToken()})
}
