package controllers

import (
	// "encoding/json"
	"log"

	// "cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	// "fmt"
	"auth-service/database"
	"auth-service/models"
	pb "auth-service/proto_buffers/protos"
)

const (
	signUpFlow = "SIGNUP"
	loginFlow = "LOGIN"
)


func(s *Server) VerifyUser(ctx context.Context, data *pb.VerifyPhoneNumberData) (*pb.VerifyPhoneNumberResponse, error) {
	log.Printf("Received request params: %v %v %v", data.PhoneNumber, data.Otp, data.Flow)
	// USER
	user, err := models.FindUser(data.PhoneNumber)
	if err != nil {
		return &pb.VerifyPhoneNumberResponse{
			Status: "Failed to get user",
		}, nil
	}

	// VERIFICATION
	verification, err := models.FindVerification(int(user.ID))
	if err != nil {
		return &pb.VerifyPhoneNumberResponse{
			Status: "Failed to find verification",
		}, nil
	}

	if verification.OtpVerified == true {
		return &pb.VerifyPhoneNumberResponse{
			Status: "No new verification present",
		}, nil
	}

	if verification.Otp != data.Otp {
		return &pb.VerifyPhoneNumberResponse{
			Status: "Otp did not match",
		}, nil
	}

	if data.Flow == signUpFlow {
		markSignedUp(database.DB, user, verification)
	} else {
		markLoggedIn(database.DB, user, verification)
	}

	return &pb.VerifyPhoneNumberResponse{
		Status: "Otp verified successfully",
	}, nil
}

func markSignedUp(db *gorm.DB, user *models.User, verification * models.Verification) {
	user.Verified = true
	verification.OtpVerified = true

	db.Save(&user)
	db.Save(&verification)
}

func markLoggedIn(db *gorm.DB, user *models.User, verification * models.Verification) {
	user.LoggedIn = true
	verification.OtpVerified = true
	
	db.Save(&user)
	db.Save(&verification)
}
