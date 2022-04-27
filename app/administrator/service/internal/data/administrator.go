package data

import (
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/model"
	"github.com/ZQCard/kratos-base-project/pkg/gormHelper"
	"github.com/ZQCard/kratos-base-project/pkg/utils/encryption"
	"github.com/ZQCard/kratos-base-project/pkg/utils/timeHelper"
	"github.com/go-kratos/kratos/v2/errors"
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

func (a administratorRepo) VerifyPassword(ctx context.Context, id int64, password string) (bool, error) {
	administrator := model.Administrator{}
	if err := a.data.db.Model(&model.Administrator{}).Where("id = ?", id).First(&administrator).Error; err != nil {
		return false, err
	}
	return encryption.CheckPassword(administrator.Password, administrator.Salt, password), nil
}

func (a administratorRepo) GetAdministrator(ctx context.Context, params map[string]interface{}) (*biz.Administrator, error) {
	if len(params) == 0 {
		return nil, errors.BadRequest("params must not be empty", "参数不能为空")
	}
	fmt.Println(params)
	response := &biz.Administrator{}
	administrator := model.Administrator{}

	// 查询db
	db := a.data.db.Model(&model.Administrator{})
	if id, ok := params["id"]; ok {
		db.Where("id = ?", id)
	}
	if username, ok := params["username"]; ok{
		db.Where("username = ?", username)
	}
	// 如果存在筛选删除 需要设定， 否则必须为未删除数据
	if hasDeleted, ok := params["has_deleted"]; ok && hasDeleted.(int64) != 0 {
		db.Scopes(gormHelper.RecordDeleted())
	}else {
		db.Scopes(gormHelper.RecordUnDeleted())
	}

	if err := db.First(&administrator).Error; err != nil {
		return nil, err
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

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &administratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
