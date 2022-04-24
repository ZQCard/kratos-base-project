package data

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/model"
	"github.com/ZQCard/kratos-base-project/pkg/errors/systemError"
	"github.com/ZQCard/kratos-base-project/pkg/utils/timeHelper"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/biz"
)


var administratorCacheKey = func(username string) string {
	return "administrator_cache_key_" + username
}

type administratorRepo struct {
	data *Data
	log  *log.Helper
}
func (a administratorRepo) FindAdministratorByUsername(ctx context.Context, username string) (*biz.Administrator, error) {
	response := &biz.Administrator{}

	administrator := model.Administrator{}
	if err := a.data.db.Model(&model.Administrator{}).Where("username = ?", username).First(&administrator).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, systemError.SystemError
	}

	response.Id = administrator.Id
	response.Username = administrator.Username
	response.Password = administrator.Password
	response.Mobile = administrator.Mobile
	response.Nickname = administrator.Nickname
	response.Avatar = administrator.Avatar
	response.Status = administrator.Status
	response.CreatedAt = timeHelper.FormatTimeYMGHIS(*administrator.CreatedAt)
	response.UpdatedAt = timeHelper.FormatTimeYMGHIS(*administrator.UpdatedAt)
	response.DeletedAt = timeHelper.FormatTimeInt64YMGHIS(administrator.DeletedAt)
	return response, nil
}

func (a administratorRepo) VerifyPassword(ctx context.Context, id int64, password string) error {
	panic("implement me")
}

func (a administratorRepo) GetAdministrator(ctx context.Context, id int64, administrator biz.Administrator) (*biz.Administrator, error) {
	panic("implement me")
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &administratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
