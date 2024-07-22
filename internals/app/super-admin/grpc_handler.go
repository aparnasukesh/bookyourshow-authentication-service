package superadmin

import (
	"context"

	"github.com/aparnasukesh/inter-communication/auth"
)

type GrpcHandler struct {
	svc Service
	auth.UnimplementedSuperAdminServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) SuperAdminAuthRequired(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	err := h.svc.SuperAdminAuthentication(ctx, req.Token)
	if err != nil {
		return &auth.AuthResponse{
			Status:     "super-admin authentication failed",
			StatusCode: 401,
		}, err
	}
	return &auth.AuthResponse{
		Status:     "super-admin authentication successfull",
		StatusCode: 200,
	}, nil
}
