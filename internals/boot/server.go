package boot

import (
	"log"
	"net"

	"github.com/aparnasukesh/auth-svc/config"
	"github.com/aparnasukesh/auth-svc/internals/app/admin"
	"github.com/aparnasukesh/auth-svc/internals/app/jwt"
	superadmin "github.com/aparnasukesh/auth-svc/internals/app/super-admin"
	"github.com/aparnasukesh/auth-svc/internals/app/user"
	pb "github.com/aparnasukesh/inter-communication/auth"
	"google.golang.org/grpc"
)

func NewGrpcServer(config config.Config, grpcHandler jwt.GrpcHandler, userHandler user.GrpcHandler, superAdminHandler superadmin.GrpcHandler, adminHandler admin.GrpcHandler) (func() error, error) {
	//lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	lis, err := net.Listen("tcp", "0.0.0.0:"+config.GrpcPort)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	pb.RegisterJWT_TokenServiceServer(s, &grpcHandler)
	pb.RegisterUserAuthServiceServer(s, &userHandler)
	pb.RegisterSuperAdminAuthServiceServer(s, &superAdminHandler)
	pb.RegisterAdminAuthServiceServer(s, &adminHandler)
	srv := func() error {
		log.Printf("gRPC server started on port %s", config.GrpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}
		return nil
	}
	return srv, nil
}
