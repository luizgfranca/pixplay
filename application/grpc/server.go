package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/luizgfranca/pixplay/application/grpc/pb"
	"github.com/luizgfranca/pixplay/application/usecase"
	"github.com/luizgfranca/pixplay/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(db *gorm.DB, port int) {
	server := grpc.NewServer()
	reflection.Register(server)

	pixRepository := repository.PixKeyRepositoryDB{Db: db}
	pixUseCase := usecase.PixKeyUseCase{PixKeyRepository: pixRepository}
	service := GetPixGRPCService(pixUseCase)
	pb.RegisterPixServiceServer(server, service)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("error starting listener", err)
	}

	log.Printf("Server listening to %s", address)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("error starting GRPC server", err)
	}
}
