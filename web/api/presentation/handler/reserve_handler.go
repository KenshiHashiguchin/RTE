package handler

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IReserveHandler interface {
	HandleReserve(c *gin.Context)
}

type reserveHandler struct {
	usecase usecase.IAuth
}

func newReserveHandler() *AuthHandler {
	dao := kvs.NewUserToken()
	userRepo := repositoryImpl.NewUserRepository()
	return &AuthHandler{
		Usecase: usecase.NewAuthUseCase(dao, userRepo),
	}
}

func HandleReserve(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{})
}
