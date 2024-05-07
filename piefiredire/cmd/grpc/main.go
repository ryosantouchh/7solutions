package main

import (
	"log"
	"net"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/adapter/storage/repository"
	beef_grpc "github.com/ryosantouchh/7solutions/piefiredire/services/grpc/beef"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("gRPC server listening failed: %v \n", err)
	}

	var mockDB map[string]interface{}

	beefRepository := repository.NewBeefRepository(mockDB)
	beefServer := beef_grpc.NewBeefgRPCServer(beefRepository)
	server := grpc.NewServer()
	beef_grpc.RegisterBeefServiceServer(server, beefServer)

	reflection.Register(server)

	log.Println("gRPC server listen on port 50051")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
