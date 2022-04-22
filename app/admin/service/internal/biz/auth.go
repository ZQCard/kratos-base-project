package biz

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"

	v1 "github.com/ZQCard/kratos-base-project/api/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/admin/service/internal/conf"
)

type AuthUseCase struct {
	key               string
	administratorRepo AdministratorRepo
}

func NewAuthUseCase(conf *conf.Auth, administratorRepo AdministratorRepo) *AuthUseCase {
	return &AuthUseCase{
		key:               conf.ApiKey,
		administratorRepo: administratorRepo,
	}
}

func (receiver *AuthUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {

	// 获取用户
	user, err := receiver.administratorRepo.FindAdministratorByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	// 验证密码
	err = receiver.administratorRepo.VerifyPassword(ctx, user.Id, req.Password)
	if err != nil {
		return nil,  err
	}
	// 生成token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(receiver.key))
	if err != nil {
		return nil, errors.New("error")
	}
	return &v1.LoginReply{
		Token: signedString,
	}, nil
}
