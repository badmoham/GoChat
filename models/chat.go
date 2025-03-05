package models

import "time"

type Message struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type uint   `gorm:"not null" json:"type"` // used to determine message type for later use. e.g:1:text 2:image ...
	Text string `json:"text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// relations
	UserID   uint `json:"user_id"`
	RoomType uint `json:"room_type"`
	RoomID   uint `json:"room_id"`
}

type P2P struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Messages  []Message `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Users     []User    `gorm:"many2many:user_p2p_chats;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
