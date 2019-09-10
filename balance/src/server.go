package main

import (
	"fmt"
	"net"

	"./proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
)

type server struct {
	Database *gorm.DB
}

func main() {
	fmt.Printf("[Client] Balance service starting...\n")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/balance?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	server := &server{
		Database: db,
	}
	proto.RegisterCommunicationServer(srv, server)

	if !db.HasTable(&AccountsModel{}) {
		db.CreateTable(&AccountsModel{})
	}

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}