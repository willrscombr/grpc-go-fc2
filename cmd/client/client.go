package main

import (
	"context"
	"fmt"
	"log"

	"github.com/codeedu/fc2-rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Não foi possível conectar no servidor: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	Adduser(client)
}

func Adduser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "12",
		Name:  "joao",
		Email: "j@j.com.br",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Não foi possível conectar no servidor: %v", err)
	}

	fmt.Println(res)
}
