package service

import (
	"context"

	pb "github.com/ZQCard/kratos-base-project/api/administrator/v1"
)


func (s *AdministratorService) GetAdministratorByUsername(ctx context.Context, req *pb.GetAdministratorByUsernameRequest) (*pb.GetAdministratorByUsernameReply, error) {
	return s.administratorCase.FindAdministratorByUsername(ctx, req)
}
func (s *AdministratorService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	return &pb.VerifyPasswordReply{}, nil
}
func (s *AdministratorService) GetAdministrator(ctx context.Context, req *pb.GetAdministratorRequest) (*pb.GetAdministratorReply, error) {
	return &pb.GetAdministratorReply{}, nil
}
