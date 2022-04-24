package data

import (
	"context"
	"errors"
	"fmt"
	administratorClientV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	"github.com/ZQCard/kratos-base-project/pkg/errors/systemError"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"

	"github.com/ZQCard/kratos-base-project/app/admin/service/internal/biz"
)

type administratorRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (rp administratorRepo) FindAdministratorByUsername(ctx context.Context, username string) (*biz.Administrator, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("find_user_by_name_%s", username), func() (interface{}, error) {
		user, err := rp.data.administratorClient.GetAdministratorByUsername(ctx, &administratorClientV1.GetAdministratorByUsernameRequest{
			Username: username,
		})
		if err != nil {
			return nil, err
		}
		return &biz.Administrator{
			Id:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Administrator), nil
}

func (rp administratorRepo) VerifyPassword(ctx context.Context, id int64, password string) error {
		reply, err := rp.data.administratorClient.VerifyPassword(ctx, &administratorClientV1.VerifyPasswordRequest{
			Id: id,
			Password: password,
		})
		if err != nil {
			rp.log.Error("systemError error : administrator service error : "+err.Error() )
			return systemError.SystemError
		}

		if reply.Success == false {
			return errors.New("密码错误")
		}
		return nil
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &administratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}
