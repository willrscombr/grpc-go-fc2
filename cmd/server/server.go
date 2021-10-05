package main

import (
	"fmt"
	"log"
	"net"

	"github.com/codeedu/fc2-rpc/pb"
	"github.com/codeedu/fc2-rpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", "localhost:50051")

	if err != nil {

		log.Fatalf("Não foi possível conectar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}

	fmt.Println("Servidor Executando")

}
