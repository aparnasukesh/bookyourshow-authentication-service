package jwt

import (
	"context"

	"github.com/aparnasukesh/auth-svc/pkg/common"
	pb "github.com/aparnasukesh/inter-communication/auth"
)

type GrpcHandler struct {
	svc common.JWT_Service
	pb.UnimplementedJWT_TokenServiceServer
}

func NewGrpcHandler(svc common.JWT_Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) GenerateJWt(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	token, err := h.svc.GenerateJWT(req.Email, uint(req.UserId), uint(req.RoleId))
	if err != nil {
		return nil, err
	}
	return &pb.GenerateResponse{
		Token: token,
	}, nil
}
