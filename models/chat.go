package models

import "time"

type Message struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type uint   `gorm:"not null" json:"type"` // used to determine message type for later use. e.g:1:text 2:image ...
	Text string `json:"text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// relations
	UserID     uint `json:"user_id"`
	ChatRoomID uint `json:"chat_room_id"`
}

type ChatRoom struct {
	ID   uint `gorm:"primaryKey" json:"id"`
	Type uint `gorm:"not null" json:"type"`

	Messages  []Message `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Users     []User    `gorm:"many2many:user_chat_room;"`
	Speaker   []User    `gorm:"many2many:speaker_chat_room;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
