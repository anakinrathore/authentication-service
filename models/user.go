package models

import (
	"auth-service/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar;"`
	PhoneNumber string `gorm:"type:varchar;"`
	Verified bool `gorm:"type:boolean"`
	LoggedIn bool `gorm:"type:boolean"`
}


func(u *User)Create() (error) {
	err := database.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func(u *User)MarkVerified() (error) {
	u.Verified=true
	database.DB.Save(&u)
	err := database.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func FindUser(phoneNumber string) (*User, error) {
	user := User{}
	err := database.DB.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

