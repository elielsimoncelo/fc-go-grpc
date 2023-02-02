package main

import (
	"database/sql"
	"net"

	"github.com/elielsimoncelo/fc-go-grpc/internal/database"
	"github.com/elielsimoncelo/fc-go-grpc/internal/pb"
	"github.com/elielsimoncelo/fc-go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server)

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
