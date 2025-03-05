package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	PhoneNumber string    `gorm:"unique not null" json:"phone_number"`
	Password    string    `gorm:"not null" json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// relations
	Messages []Message `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	P2Ps     []P2P     `gorm:"many2many:user_p2p_chats;"`
}
