package controllers

import (
	"log"
	"golang.org/x/net/context"
	"auth-service/models"
	pb "auth-service/proto_buffers/protos"
)



func(s *Server) GetUserProfile(ctx context.Context, data *pb.GetUserProfileData) (*pb.GetUserProfileResponse, error) {
	log.Printf("Received request params: %v", data.PhoneNumber)
	user, err := models.FindUser(data.PhoneNumber)
	if err != nil || user.LoggedIn == false {
		return &pb.GetUserProfileResponse{
			Name: "",
			PhoneNumber: "",
		}, nil
	}

	return &pb.GetUserProfileResponse{
		Name: user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
