package admin

import (
	"context"

	"github.com/aparnasukesh/inter-communication/auth"
)

type GrpcHandler struct {
	svc Service
	auth.UnimplementedAdminAuthServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) AdminAuthRequired(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	err := h.svc.AdminAuthentication(ctx, req.Token)
	if err != nil {
		return &auth.AuthResponse{
			Status:     "admin authentication failed",
			StatusCode: 401,
		}, err
	}
	return &auth.AuthResponse{
		Status:     "admin authentication successfull",
		StatusCode: 200,
	}, nil
}
