package services

import (
	"errors"

	"GoChat/config"
	"GoChat/enums"
	"GoChat/models"
)

func CreateNewPersonalChat(sourceUserID, destUserID uint) (uint, error) {
	searchQuery := config.DB.Table("user_chat_room").
		Joins("JOIN chat_rooms ON chat_rooms.id = user_chat_room.chat_room_id").
		Where("user_chat_room.user_id IN (?, ?)", sourceUserID, destUserID).
		Where("chat_rooms.type = ?", enums.ChatP2P).
		Group("chat_rooms.id").
		Having("COUNT(DISTINCT user_chat_room.user_id) = 2")

	var exists bool
	err := config.DB.Table("user_chat_room").
		Select("1").
		Where("EXISTS (?)", searchQuery).
		Limit(1). // Limit to 1 result for efficiency
		Scan(&exists).Error

	if err != nil {
		return 0, errors.New("server query failed")
	}
	if exists {
		var chatRoomID uint
		err := searchQuery.Select("chat_room.id").Scan(&chatRoomID).Error
		if err != nil {
			return 0, err
		}
		return chatRoomID, nil // room already exists, so we will just return it without creating a new
	}

	// this transaction will be atomic to prevent creating room without its user
	tx := config.DB.Begin()
	if tx.Error != nil {
		return 0, errors.New("transaction failed")
	}
	// Rollback in case of any error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback if a panic occurs
		}
	}()

	// Create the ChatRoom
	chatRoom := models.ChatRoom{
		Type: 1,
	}
	if err := tx.Create(&chatRoom).Error; err != nil {
		tx.Rollback() // Rollback if creating the ChatRoom fails
		return 0, err
	}

	var chatAttendants = []models.User{{ID: sourceUserID}, {ID: destUserID}}
	if err := tx.Model(&chatRoom).Association("Users").Append(chatAttendants); err != nil {
		tx.Rollback() // Rollback if associating Users fails
		return 0, err
	}
	if err := tx.Model(&chatRoom).Association("Speaker").Append(chatAttendants); err != nil {
		tx.Rollback() // Rollback if associating Users fails
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return chatRoom.ID, nil
}

func GetUserAllChats(UserID uint) ([]uint, error) {
	var chatRoomIDs []uint
	err := config.DB.Table("user_chat_room").
		Joins("JOIN chat_rooms ON chat_rooms.id = user_chat_room.chat_room_id").
		Where("user_chat_room.user_id = ?", UserID).
		Select("chat_rooms.id").
		Scan(&chatRoomIDs).Error
	if err != nil {
		return []uint{}, err
	}
	return chatRoomIDs, nil
}
