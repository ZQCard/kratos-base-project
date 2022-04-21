package data

import (
	administratorClientV1 "admin/api/administrator/v1"
	"admin/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

type administratorRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (rp administratorRepo) FindAdministratorByUsername(ctx context.Context, username string) (*biz.Administrator, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("find_user_by_name_%s", username), func() (interface{}, error) {
		fmt.Println(rp.data.administratorClient)
		user, err := rp.data.administratorClient.GetAdministratorByUsername(ctx, &administratorClientV1.GetAdministratorByUsernameRequest{
			Username: username,
		})
		if err != nil {
			return nil, biz.ErrAdministratorNotFound
		}
		return &biz.Administrator{
			ID:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Administrator), nil
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &administratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}
