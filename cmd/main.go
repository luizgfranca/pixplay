package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/luizgfranca/pixplay/application/grpc"
	"github.com/luizgfranca/pixplay/infra/db"
)

var database *gorm.DB

func main() {
	fmt.Println("starting server")
	database := db.ConnectDB(os.Getenv("env"))
	grpc.StartGRPCServer(database, 50051)
}
