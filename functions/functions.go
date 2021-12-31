package funtions

import (
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

//Screenshot function will save the screenshot and send it.
func Screenshot(update tgbotapi.Update) {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}

	bot, error := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if error != nil {
		panic(error)
	}

	website := update.Message.Text

	if !strings.Contains(website, "http") {
		website = "http://" + website
	}

	if !strings.Contains(website, ".") {
		website = website + ".com"
	}

	url := "https://api.screenshotmachine.com/?key=" + getRandomKey() + "&url=" + website + "&dimension=1024x3080"

	screenshot := tgbotapi.NewMessage(update.Message.Chat.ID, url)
	bot.Send(screenshot)
}

//get getRandomKey will return a random item from a slice.
func getRandomKey() string {
	items := []string{"9cf822", "764f21", "6244bf ", "97ec9a", "80e53c", "f2909b", "7aaf56", "e7f994"}
	return items[rand.Intn(len(items))]
}
