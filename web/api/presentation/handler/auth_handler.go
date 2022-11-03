package handler

import (
	"errors"
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/presentation/request"
	"github.com/RTE/web/api/usecase"
	"github.com/RTE/web/api/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	Usecase usecase.IAuth
}

func newAuthHandler() *AuthHandler {
	dao := kvs.NewUserToken()
	return &AuthHandler{
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

func AuthUser(c *gin.Context) (model.UserClaims, error) {
	session := sessions.Default(c)
	authSession := session.Get("auth")

	if authSession == nil {
		return model.UserClaims{}, errors.New("no user session")
	}

	return model.GetAuthUserByTokenString(authSession.(string), util.Env("JWT_SECRET_KEY"))
}
