package service

import (
	v1 "admin/api/admin/v1"
	"context"
)


func (s *AdminInterface) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	s.authCase.
	return &v1.LoginReply{}, nil
}
func (s *AdminInterface) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}