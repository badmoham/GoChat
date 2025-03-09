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
	Messages       []Message  `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	ChatRooms      []ChatRoom `gorm:"many2many:user_chat_room;"`
	SpeakableRooms []ChatRoom `gorm:"many2many:speaker_chat_room;"`
}
