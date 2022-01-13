package controllers

import (
	"log"

	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"auth-service/models"
	"auth-service/utilities"
	pb "auth-service/proto_buffers/protos"
)

func(s *Server) LoginUser(ctx context.Context, data *pb.LoginData) (*pb.LoginResponse, error) {
	log.Printf("Received request params: %v", data.PhoneNumber)
	user, err := models.FindUser(data.PhoneNumber)
	if err != nil {
		return &pb.LoginResponse{
			Status: "Failed to get user",
		}, nil
	}

	if user.LoggedIn == true{
		return &pb.LoginResponse{
			Status: "User already signed in",
		}, nil
	}

	otp := utilities.GenerateOTP()

	verification := models.Verification{
		UserId: int(user.ID),
		OtpVerified: false,
		Otp: otp,
	}
	
	e := verification.Create()
	if e != nil {
		return &pb.LoginResponse{
			Status: "Verification Creation Failed",
		}, nil
	}

	msgUuid := uuid.NewV4().String()

	eventData := utilities.EventData{
		ID: msgUuid,
		Name: user.Name,
		PhoneNumber: user.PhoneNumber,
     	Otp: otp, 
	}
	err = utilities.Publish(ctx, eventData, otp)
	
	if err != nil {
		return &pb.LoginResponse{
			Status: err.Error(),
		}, nil
	}
	return &pb.LoginResponse{
		Status: "User login otp sent",
	}, nil
}
