package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	PhoneNumber string    `gorm:"unique not null" json:"phone_number"`
	Password    string    `gorm:"not null" json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
