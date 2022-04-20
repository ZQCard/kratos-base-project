package service

import (
	v1 "admin/api/admin/v1"
	"admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
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
