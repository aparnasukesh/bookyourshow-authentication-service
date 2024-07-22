package superadmin

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aparnasukesh/auth-svc/pkg/common"
)

type service struct {
	svc common.JWT_Service
}

type Service interface {
	SuperAdminAuthentication(ctx context.Context, authorization string) error
}

func NewSuperAdminService(svc common.JWT_Service) Service {
	return &service{
		svc: svc,
	}
}

func (s *service) SuperAdminAuthentication(ctx context.Context, authorization string) error {
	if authorization == "" {
		return errors.New("authorization header is missing")
	}

	tokenParts := strings.Split(authorization, "Bearer ")
	if len(tokenParts) < 2 {
		return errors.New("Bearer token is missing or malformed. Ensure your Authorization header is in the format 'Bearer <token>'")
	}

	verifiedToken, err := s.svc.VerifyJWT(tokenParts[1])
	if err != nil {
		return fmt.Errorf("failed to verify JWT token: %w", err)
	}

	roleId, err := s.svc.GetRole(verifiedToken)
	if err != nil {
		return fmt.Errorf("failed to get role from token: %w", err)
	}

	if roleId != 3 {
		return errors.New("access denied: you do not have the required role to perform this action")
	}
	return nil
}
