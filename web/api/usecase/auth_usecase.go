package usecase

import (
	"errors"
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/domain/repository"
	"github.com/RTE/web/api/infrastructure/dao/kvs"
	"github.com/RTE/web/api/util"
	"github.com/dgrijalva/jwt-go"
)

type IAuth interface {
	Auth(address string, signature string) (string, error)
}

type auth struct {
	userTokenDao   kvs.IUserToken
	userRepository repository.User
}

func NewAuthUseCase(userTokenDao kvs.IUserToken, userRepository repository.User) IAuth {
	return &auth{
		userTokenDao:   userTokenDao,
		userRepository: userRepository,
	}
}

func (r auth) Auth(address string, signature string) (string, error) {
	// verify signature
	token := r.userTokenDao.GetToken(address)
	if len(token) == 0 {
		return "", errors.New("A token that does not exist.")
	}

	userSig := model.NewUserSignature(address, signature, token)
	err := userSig.Verify()
	if err != nil {
		return "", err
	}

	claims := model.NewUserClaims(address)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := jwtToken.SignedString([]byte(util.Env("JWT_SECRET_KEY")))

	return tokenString, nil
}
