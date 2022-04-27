package biz

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/model"
	"github.com/ZQCard/kratos-base-project/pkg/errors/administratorError"
	"github.com/ZQCard/kratos-base-project/pkg/errors/systemError"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

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
	VerifyPassword(ctx context.Context, id int64, password string) (bool, error)
	GetAdministrator(ctx context.Context, params map[string]interface{})  (*Administrator, error)
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

func (ac AdministratorUseCase)FindLoginAdministratorByUsername(ctx context.Context, in *v1.GetLoginAdministratorByUsernameRequest) (*v1.GetLoginAdministratorByUsernameReply, error) {
	params := make(map[string]interface{})
	params["username"] = in.Username
	administrator, err := ac.repo.GetAdministrator(ctx, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &v1.GetLoginAdministratorByUsernameReply{}, administratorError.AdministratorNotExist
		}
		return &v1.GetLoginAdministratorByUsernameReply{}, systemError.SystemError
	}
	if administrator.Status == model.AdministratorStatusForbid {
		return &v1.GetLoginAdministratorByUsernameReply{}, administratorError.AdministratorForbid
	}
	if administrator.DeletedAt != "" {
		return &v1.GetLoginAdministratorByUsernameReply{}, administratorError.AdministratorDeleted
	}

	return &v1.GetLoginAdministratorByUsernameReply{
		Id:        administrator.Id,
		Username:  administrator.Username,
	}, nil
}


func (ac AdministratorUseCase)FindAdministratorById(ctx context.Context, id int64) (*v1.GetAdministratorReply, error) {
	params := make(map[string]interface{})
	params["id"] = id
	administrator, err := ac.repo.GetAdministrator(ctx, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &v1.GetAdministratorReply{}, administratorError.AdministratorNotExist
		}
		return &v1.GetAdministratorReply{}, systemError.SystemError
	}
	if administrator.Status == model.AdministratorStatusForbid {
		return &v1.GetAdministratorReply{}, administratorError.AdministratorForbid
	}
	if administrator.DeletedAt != "" {
		return &v1.GetAdministratorReply{}, administratorError.AdministratorDeleted
	}

	return &v1.GetAdministratorReply{
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
}



func (ac AdministratorUseCase)VerifyAdministratorPassword(ctx context.Context, in *v1.VerifyPasswordRequest) (bool, error) {
	result, err := ac.repo.VerifyPassword(ctx, in.Id, in.Password)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, systemError.SystemError
	}
	return result, nil
}