package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGRPCServer(db *gorm.DB, port int) {
	server := grpc.NewServer()
	address := fmt.Sprintf("0.0.0.0:#{port}")
	listener, err := net.Listen("tpc", address)
	if err != nil {
		log.Fatal("error starting listener", err)
	}

	log.Printf("Server listening to %s", address)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("error starting GRPC server", err)
	}

}
