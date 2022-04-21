package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrAdministratorNotFound = errors.New("user not found")
)

type Administrator struct {
	ID       int64
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
