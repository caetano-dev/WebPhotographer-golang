package funtions

import (
	"encoding/json"
	"io/ioutil"
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

	url := "https://api.screenshotmachine.com/?key=" + randomKey() + "&url=" + website + "&dimension=1024x3080"

	screenshot := tgbotapi.NewMessage(update.Message.Chat.ID, url)
	bot.Send(screenshot)
}

// Keys struct will hold the keys.
type Keys struct {
	Keys []string `json:"keys"`
}

// randomKey will return a random item from a slice.
func randomKey() string {
	file, error := os.Open("keys.json")
	if error != nil {
		log.Fatal(error)
	}

	keys, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatal(error)
	}

	var keysStruct Keys
	json.Unmarshal(keys, &keysStruct)

	randomKey := keysStruct.Keys[rand.Intn(len(keysStruct.Keys))]

	return randomKey

}
