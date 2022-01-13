package controllers

import (
	"log"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"auth-service/models"
	"auth-service/utilities"
	pb "auth-service/proto_buffers/protos"
)

type Server struct {
	
}

func(s *Server) CreateUser(ctx context.Context, data *pb.SignUpData) (*pb.SignUpResponse, error) {
	log.Printf("Received request params: %v %v", data.PhoneNumber, data.Name)
	user := models.User{
		Name: data.Name,
		PhoneNumber: data.PhoneNumber,
		Verified: false,
		LoggedIn: false,
	}
	e := user.Create()
	if e != nil {
		return &pb.SignUpResponse{
			Status: "User Creation Failed",
		}, nil
	}
	otp := utilities.GenerateOTP()

	verification := models.Verification{
		UserId: int(user.ID),
		OtpVerified: false,
		Otp: otp,
	}

	e = verification.Create()
	if e != nil {
		return &pb.SignUpResponse{
			Status: "Verification Creation Failed",
		}, nil
	}

	msgUuid := uuid.NewV4().String()
	eventData := utilities.EventData{
		ID: msgUuid,
		Name: data.Name,
		PhoneNumber: data.PhoneNumber,
     	Otp: otp, 
	}
	err := utilities.Publish(ctx, eventData, otp)
	
	if err != nil {
		return &pb.SignUpResponse{
			Status: err.Error(),
		}, nil
	}
	return &pb.SignUpResponse{
		Status: "User signup otp sent",
	}, nil
}


