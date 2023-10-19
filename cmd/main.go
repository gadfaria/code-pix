package main

import (
	"os"

	"github.com/gadfaria/code-pix/application/grpc"
	"github.com/gadfaria/code-pix/infrastructure/db"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
