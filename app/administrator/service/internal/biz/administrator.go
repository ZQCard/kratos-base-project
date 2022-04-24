package biz

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/model"
	"github.com/ZQCard/kratos-base-project/pkg/errors/administratorError"
	"github.com/ZQCard/kratos-base-project/pkg/errors/systemError"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
)

type Administrator struct {
	Id       int64
	Username string
	Password string
	Mobile   string
	Nickname string
	Avatar   string
	Status   int64
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type AdministratorRepo interface {
	FindAdministratorByUsername(ctx context.Context, username string) (*Administrator, error)
	VerifyPassword(ctx context.Context, id int64, password string) error
	GetAdministrator(ctx context.Context, id int64,administrator Administrator)  (*Administrator, error)
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

func (ac AdministratorUseCase)FindAdministratorByUsername(ctx context.Context, in *v1.GetAdministratorByUsernameRequest) (*v1.GetAdministratorByUsernameReply, error) {
	administrator, err := ac.repo.FindAdministratorByUsername(ctx, in.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &v1.GetAdministratorByUsernameReply{}, administratorError.AdministratorNotExist
	}
	if err != nil {
		return &v1.GetAdministratorByUsernameReply{}, systemError.SystemError
	}
	if administrator.Status == model.AdministratorStatusForbid {
		return &v1.GetAdministratorByUsernameReply{}, administratorError.AdministratorForbid
	}
	if administrator.DeletedAt != "" {
		return &v1.GetAdministratorByUsernameReply{}, administratorError.AdministratorDeleted
	}

	return &v1.GetAdministratorByUsernameReply{
		Id:        administrator.Id,
		Username:  administrator.Username,
		Password:  administrator.Password,
		Mobile:    administrator.Mobile,
		Nickname:  administrator.Nickname,
		Avatar:    administrator.Avatar,
		Status:    administrator.Status,
		CreatedAt: administrator.CreatedAt,
		UpdatedAt: administrator.UpdatedAt,
		DeletedAt: administrator.DeletedAt,
	}, nil
	return nil, nil
}