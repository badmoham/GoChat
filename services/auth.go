package services

import (
	"fmt"
	"time"

	"GoChat/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	//will compute password hash with default computational cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func CheckPassword(hashedPassword, password string) error {
	// if password with its hash does not match, will return an error
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EncryptJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userID, "exp": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func DecryptJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("token incalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("token is invalid")
	}
	userID, ok := claims["sub"].(uint)
	if !ok {
		return 0, fmt.Errorf("token invalid")
	}
	return userID, nil
}
