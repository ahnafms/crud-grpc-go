package main

import (
	// DB "github.com/ahnafms/pi-grpc-go/pkg/database"
	"log"
  "net"
	"google.golang.org/grpc"
  moviePb "github.com/ahnafms/crud-grpc-go/pkg/protobuf"
  "github.com/ahnafms/crud-grpc-go/cmd/services"
  DB "github.com/ahnafms/crud-grpc-go/pkg/database"  
)

const (
  port = ":50051"
)

func main(){ 
  DB.ConnectDB()
  netListen, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("Failed to listen %v", err.Error())
  }
  grpcServer := grpc.NewServer()
  movieService := services.MovieService{}
  moviePb.RegisterMovieServiceServer(grpcServer, &movieService)

  log.Printf("Server started at %v", netListen.Addr())
  if err := grpcServer.Serve(netListen); err != nil{
    log.Fatal("Failed to server 8080", err.Error())
  }
}

