package service

import (
	"context"

	pb "github.com/ZQCard/kratos-base-project/api/administrator/v1"
)


func (s *AdministratorService) GetLoginAdministratorByUsername(ctx context.Context, req *pb.GetLoginAdministratorByUsernameRequest) (*pb.GetLoginAdministratorByUsernameReply, error) {
	return s.administratorCase.FindLoginAdministratorByUsername(ctx, req)
}

func (s *AdministratorService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	res, err := s.administratorCase.VerifyAdministratorPassword(ctx, req)
	return &pb.VerifyPasswordReply{
		Success: res,
	}, err
}

func (s *AdministratorService) GetAdministrator(ctx context.Context, req *pb.GetAdministratorRequest) (*pb.GetAdministratorReply, error) {
	return s.administratorCase.FindAdministratorById(ctx, req.Id)
}
