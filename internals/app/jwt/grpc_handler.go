package jwt

import (
	"context"
	"fmt"

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

func (h *GrpcHandler) VerifyJWT(ctx context.Context, req *pb.VerifyJWTRequest) (*pb.VerifyJWTResponse, error) {
	verifiedToken, err := h.svc.VerifyJWT(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyJWTResponse{
		Token: verifiedToken.Raw,
	}, nil
}

func (h *GrpcHandler) GetUserID(ctx context.Context, req *pb.GetUserIDRequest) (*pb.GetUserIDResponse, error) {
	verifiedToken, err := h.svc.VerifyJWT(req.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to verify jwt token %w", err)
	}
	userId, err := h.svc.GetUserID(verifiedToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get userid from token %w", err)
	}
	return &pb.GetUserIDResponse{
		UserId: int32(userId),
	}, nil
}
