package user

import (
	"context"

	"github.com/aparnasukesh/inter-communication/auth"
)

type GrpcHandler struct {
	svc Service
	auth.UnimplementedUserAuthServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}
func (h *GrpcHandler) UserAuthRequired(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	token := req.Token

	err := h.svc.UserAuthentication(ctx, token)
	if err != nil {
		return &auth.AuthResponse{
			Status:     "user authentication failed",
			StatusCode: 401,
		}, err
	}
	return &auth.AuthResponse{
		Status:     "user authentication successful",
		StatusCode: 200,
	}, nil
}
