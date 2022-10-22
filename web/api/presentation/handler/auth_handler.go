package handler

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/infrastructure/repositoryImpl"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthHandler interface {
	HandleRegisterUser(c *gin.Context)
}

type authHandler struct {
	usecase usecase.IAuth
}

func newAuthHandler() *authHandler {
	dao := kvs.NewUserToken()
	userRepo := repositoryImpl.NewUserRepository()
	return &authHandler{
		usecase: usecase.NewAuthUseCase(dao, userRepo),
	}
}

func HandleAuth(c *gin.Context) {
	auth := request.Auth{Address: c.PostForm("address"), Signature: c.PostForm("signature")}
	ok, err := auth.Validate() // Authenticate
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	u := newAuthHandler().usecase
	token, authErr := u.Auth(auth.Address, auth.Signature)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": authErr.Error()})
	}

	session := sessions.Default(c)
	session.Set("auth", token)
	session.Save()

	c.JSON(http.StatusOK, gin.H{})
}
