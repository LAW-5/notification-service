package shared

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type Notification struct {
	UserId	 int	`json:"userId"`
	Header	string	`json:"header"`
	Message	string	`json:"message"`
}

func NotificationHandler(arg []byte, db *gorm.DB) {
	var message Notification
	err := json.Unmarshal(arg[:], &message)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.Create(message)

	log.Printf("%d - %s - %s", message.UserId, message.Header, message.Message)
}