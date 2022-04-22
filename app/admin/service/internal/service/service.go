package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/ZQCard/kratos-base-project/api/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/admin/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminInterface)

type AdminInterface struct {
	v1.UnimplementedAdminServer
	administratorCase *biz.AdministratorUseCase
	authCase *biz.AuthUseCase
	log *log.Helper
}

func NewAdminInterface(
	administratorCase *biz.AdministratorUseCase,
	authCase *biz.AuthUseCase,
	logger log.Logger) *AdminInterface {
	return &AdminInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorCase:  administratorCase,
		authCase:  authCase,
	}
}
