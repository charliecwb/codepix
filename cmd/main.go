package main

import (
	"github.com/charliecwb/codepix/application/grpc"
	"github.com/charliecwb/codepix/infraestructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var dataBase *gorm.DB

func main() {
	dataBase = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(dataBase, 50051)
}
