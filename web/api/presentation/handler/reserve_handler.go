package handler

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reserveHandler struct {
	usecase usecase.IAuth
}

func newReserveHandler() *AuthHandler {
	dao := kvs.NewUserToken()
	return &AuthHandler{
		Usecase: usecase.NewAuthUseCase(dao),
	}
}

func HandleReserve(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{})
}
