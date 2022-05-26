package main

import (
	"log"
	"notification/database"
	"notification/shared"
	"notification/utils"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	f, err := os.OpenFile("log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	utils.LoadConfig()

	db, err:= database.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	shared.NewNotificationGRPCServer(db)

	connectMQ, err := amqp.Dial(utils.ApiConfig.AMQPServerURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connectMQ.Close()

	channelMQ, err := connectMQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channelMQ.Close()

	messages, err := channelMQ.Consume(
		"Notification",
		"",
		true, 
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	forever := make(chan bool)

	go func ()  {
		for m := range messages {
			shared.NotificationHandler(m.Body, db)
		}
	}()

	<-forever
}
