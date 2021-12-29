package funtions

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

//Screenshot function will save the image in the disk and send it
func Screenshot(update tgbotapi.Update) {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")
	bot, error := tgbotapi.NewBotAPI(TOKEN)

	if error != nil {
		panic(error)
	}

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	url := update.Message.Text

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	if !strings.Contains(url, ".") {
		url = url + ".com"
	}

	filename := "screenshot.png"

	var imageBuf []byte
	if error := chromedp.Run(ctx, ScreenshotTasks(url, &imageBuf)); error != nil {
		log.Fatal(error)
	}

	if error := ioutil.WriteFile(filename, imageBuf, 0644); error != nil {
		log.Fatal(error)
	}

	photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(filename))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here it is!")
	bot.Send(msg)
	bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{photo}))

}

//ScreenshotTasks function will take a screenshot of the website
func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (error error) {
			*imageBuf, error = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			return error
		}),
	}
}
