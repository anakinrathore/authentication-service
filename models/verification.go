package models

import (
	"auth-service/database"

	"gorm.io/gorm"
)

type Verification struct {
	gorm.Model
	UserId int `gorm:"type:serial;"`
	Otp string `gorm:"type:varchar;"`
	OtpVerified bool `gorm:"type:boolean"`
}

func(v *Verification)Create() (error) {
	err := database.DB.Create(v).Error
	if err != nil {
		return err
	}
	return nil
}

func FindVerification(userId int) (*Verification, error) {
	v := Verification{}
	err := database.DB.Where("user_id = ?", userId).Order("updated_at desc").First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}