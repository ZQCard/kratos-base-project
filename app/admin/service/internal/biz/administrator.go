package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Administrator struct {
	Id       int64
	Username string
	Password string
	salt     string
	mobile   string
	nickname string
	avatar   string
	status   int64
}

type AdministratorRepo interface {
	FindAdministratorByUsername(ctx context.Context, username string) (*Administrator, error)
	VerifyPassword(ctx context.Context, id int64, password string) error
}

type AdministratorUseCase struct {
	repo AdministratorRepo
	log  *log.Helper
}

func NewAdministratorUseCase(repo AdministratorRepo, logger log.Logger) *AdministratorUseCase {
	logs := log.NewHelper(log.With(logger, "module", "administrator/interface"))
	return &AdministratorUseCase{
		repo: repo,
		log:  logs,
	}
}
