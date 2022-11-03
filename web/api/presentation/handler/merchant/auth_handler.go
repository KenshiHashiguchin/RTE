package merchant

import (
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newAuthHandler() *handler.AuthHandler {
	dao := kvs.NewMerchantToken()
	return &handler.AuthHandler{
		Usecase: usecase.NewAuthUseCase(dao),
	}
}

func HandleAuth(c *gin.Context) {
	var auth request.Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := auth.Validate() // Authenticate
	if ok != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	u := newAuthHandler().Usecase
	token, authErr := u.Auth(auth.Address, auth.Signature)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": authErr.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("auth", token)
	session.Save()

	c.JSON(http.StatusOK, gin.H{})
}
