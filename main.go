package main

import (
	"log"
	"os"

	f "telegram-bot/functions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Looking for "+update.Message.Text+"...")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

			f.HandleCommands(update)
			f.Screenshot(update)
		}

	}
}