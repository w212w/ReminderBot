package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	tgClient "github.com/w212w/ReminderBot/clients/telegram"
	eventconsumer "github.com/w212w/ReminderBot/consumer/event-consumer"
	"github.com/w212w/ReminderBot/events/telegram"
	"github.com/w212w/ReminderBot/storage/sqlite"
	_ "modernc.org/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		log.Fatal("token is not specified")
	}

	return token
}
