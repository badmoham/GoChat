package services

import (
	"errors"

	"GoChat/config"
	"GoChat/models"
)

func CreateUser(user *models.User) error {
	// will add new user to db or return err
	hashedPassord, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassord
	result := config.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AuthenticateUser(phoneNumber, password string) (string, error) {
	var user models.User

	result := config.DB.Where("phone_number = ?", phoneNumber).First(user)
	if result.Error != nil {
		return "", errors.New("user not found, neither username or password is wrong")
	}
	err := CheckPassword(user.Password, password)
	if err != nil {
		return "", errors.New("user not found, neither username or password is wrong")
	}
	tokenString, err := EncryptJWT(user.ID)
	if err != nil {
		return "", errors.New("an error occurred while returning JWT")
	}
	return tokenString, nil
}
