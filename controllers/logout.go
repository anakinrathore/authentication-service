package controllers

import (
	"log"

	"auth-service/database"
	"auth-service/models"
	"golang.org/x/net/context"
	pb "auth-service/proto_buffers/protos"
)

func(s *Server) LogoutUser(ctx context.Context, data *pb.LogoutData) (*pb.LogoutResponse, error) {
	log.Printf("Received request params: %v", data.PhoneNumber)
	user, err := models.FindUser(data.PhoneNumber)
	if err != nil {
		return &pb.LogoutResponse{
			Status: "Failed to get user",
		}, nil
	}

	if user.LoggedIn == false {
		return &pb.LogoutResponse{
			Status: "User already logged out",
		}, nil
	}

	user.LoggedIn = false
	database.DB.Save(&user)

	return &pb.LogoutResponse{
		Status: "User logged out successfully",
	}, nil
}
