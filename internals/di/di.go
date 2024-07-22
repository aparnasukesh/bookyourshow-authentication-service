package di

import (
	"log"

	"github.com/aparnasukesh/auth-svc/config"
	"github.com/aparnasukesh/auth-svc/internals/app/admin"
	"github.com/aparnasukesh/auth-svc/internals/app/jwt"
	superadmin "github.com/aparnasukesh/auth-svc/internals/app/super-admin"
	"github.com/aparnasukesh/auth-svc/internals/app/user"
	"github.com/aparnasukesh/auth-svc/internals/boot"
)

func InitResources(cfg config.Config) (func() error, error) {
	// jwt module initializing
	jwtSvc := jwt.NewJWTService(cfg.JWT_secret_key)
	jwtHandler := jwt.NewGrpcHandler(jwtSvc)

	// user module initialization
	userSvc := user.NewUserService(jwtSvc)
	userHandler := user.NewGrpcHandler(userSvc)

	// admin module initialization
	adminSvc := admin.NewAdminService(jwtSvc)
	adminHandler := admin.NewGrpcHandler(adminSvc)

	// super-adimin module initialization
	superAdminSvc := superadmin.NewSuperAdminService(jwtSvc)
	superAdminHandler := superadmin.NewGrpcHandler(superAdminSvc)

	server, err := boot.NewGrpcServer(cfg, jwtHandler, userHandler, superAdminHandler, adminHandler)
	if err != nil {
		log.Fatal(err)
	}
	return server, nil
}
