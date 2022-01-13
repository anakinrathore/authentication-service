package main

import (
	"log"
	"net"
	"auth-service/controllers"
	"google.golang.org/grpc"
	"os"
	"auth-service/database"
	pb "auth-service/proto_buffers/protos"
	
)

func main(){
	database.ConnectToDB()
	pwd, err := os.Getwd()
    if err != nil {
        os.Exit(1)
    }
	credentialFileName := "test_key.json"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", pwd+"/"+credentialFileName)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil{
		log.Fatalf("Error while opening connection on port 9000:, %v", err)
	}

	s := controllers.Server{}
	grpcSer := grpc.NewServer()
	pb.RegisterSignUpServiceServer(grpcSer, &s)
	pb.RegisterUserVerifyServiceServer(grpcSer, &s)
	pb.RegisterGetUserProfileServiceServer(grpcSer, &s)
	pb.RegisterLoginServiceServer(grpcSer, &s)
	pb.RegisterLogoutServiceServer(grpcSer, &s)

	err = grpcSer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve request on grpc server on port 9000: %v", err)
	}

}