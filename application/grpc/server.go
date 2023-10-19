package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/gadfaria/code-pix/application/grpc/pb"
	"github.com/gadfaria/code-pix/application/usecase"
	"github.com/gadfaria/code-pix/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("[GRPC] Cannot start server", err)
	}

	log.Printf("[GRPC] Server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("[GRPC] Cannot start server", err)
	}

}
