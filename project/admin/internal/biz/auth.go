package biz

import (
	"context"

	v1 "admin/api/admin/v1"
	"admin/internal/conf"


	"github.com/golang-jwt/jwt/v4"
)

type AuthUseCase struct {
	key      string
	administratorRepo AdministratorRepo
}

func NewAuthUseCase(conf *conf.Auth, administratorRepo AdministratorRepo) *AuthUseCase {
	return &AuthUseCase{
		key:      conf.ApiKey,
		administratorRepo: administratorRepo,
	}
}

func (receiver *AuthUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {

	// get user
	user, err := receiver.administratorRepo.FindAdministratorByUsername(ctx, req.Username)
	if err != nil {
		return nil, v1.ErrorLoginFailed("user not found: %s", err.Error())
	}
	// check permission(password blacklist etc...)
	err = receiver.userRepo.VerifyPassword(ctx, user, req.Password)
	if err != nil {
		return nil, v1.ErrorLoginFailed("password not match")
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	signedString, err := claims.SignedString([]byte(receiver.key))
	if err != nil {
		return nil, v1.ErrorLoginFailed("generate token failed: %s", err.Error())
	}
	return &v1.LoginReply{
		Token: signedString,
	}, nil
}