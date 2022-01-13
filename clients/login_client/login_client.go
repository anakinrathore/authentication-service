package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "auth-service/proto_buffers/protos"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	c := pb.NewLoginServiceClient(conn)

	data := pb.LoginData{
		PhoneNumber: "7742068291",
	}
	response, err := c.LoginUser(context.Background(), &data)
	if err != nil {
		log.Fatalf("Failed to call server: %v", err)
	}
	log.Printf("Login response: %v", response)
}