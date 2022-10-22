package usecase

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/RTE/web/api/infrastructure/dao/kvs"
)

type IGenerateUserToken interface {
	GenerateUserToken(address string) *model.UserToken
}

type generateUserToken struct {
	userTokenDao kvs.IUserToken
}

func NewGenerateUserTokenUsecase(userTokenDao kvs.IUserToken) IGenerateUserToken {
	return &generateUserToken{userTokenDao}
}

func (t generateUserToken) GenerateUserToken(address string) *model.UserToken {
	token := model.NewUserToken(address)
	// 保存
	t.userTokenDao.SetToken(token.GetAddress(), token.GetToken())
	return token
}
