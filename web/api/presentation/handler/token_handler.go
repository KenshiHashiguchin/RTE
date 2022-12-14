package handler

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserToken struct {
	UseCase usecase.IGenerateUserToken
}

func newUserToken() *UserToken {
	dao := kvs.NewUserToken()
	return &UserToken{
		UseCase: usecase.NewGenerateUserTokenUsecase(dao),
	}
}

func HandleGetUserToken(c *gin.Context) {
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
