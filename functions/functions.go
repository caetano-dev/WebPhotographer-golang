package funtions

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	url string
)

// Keys struct will hold the keys.
type Keys struct {
	Keys []string `json:"keys"`
}

//Screenshot function will save the screenshot and send it.
func Screenshot(update tgbotapi.Update) {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}

	bot, error := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	checkError(error)

	website := update.Message.Text

	if !strings.Contains(website, "http") {
		website = "http://" + website
	}

	if !strings.Contains(website, ".") {
		website = website + ".com"
	}

	url = "https://api.screenshotmachine.com/?key=" + randomKey() + "&url=" + website + "&dimension=1024x3080"
	file, err := os.Create("screenshot.png")

	checkError(err)

	putFile(file, httpClient())

	photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath("./screenshot.png"))
	bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{photo}))
}

// randomKey will return a random item from a slice.
func randomKey() string {
	file, error := os.Open("keys.json")
	checkError(error)

	keys, error := ioutil.ReadAll(file)
	checkError(error)

	var keysStruct Keys
	json.Unmarshal(keys, &keysStruct)

	randomKey := keysStruct.Keys[rand.Intn(len(keysStruct.Keys))]

	return randomKey

}
func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(url)

	checkError(err)

	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
