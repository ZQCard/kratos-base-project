package service

import (
	"context"

	v1 "github.com/ZQCard/kratos-base-project/api/admin/v1"
)

func (s *AdminInterface) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.authCase.Login(ctx, req)
}
func (s *AdminInterface) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}
