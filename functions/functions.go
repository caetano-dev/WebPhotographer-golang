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
	if error := chromedp.Run(ctx, screenshotTasks(url, &imageBuf)); error != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Website not found.")
		bot.Send(msg)
	}

	if error := ioutil.WriteFile(filename, imageBuf, 0644); error != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "An error happened, please try again.")
		bot.Send(msg)
	}

	photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(filename))
	bot.SendMediaGroup(tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{photo}))

}

//screenshotTasks will navigate to the website and take a screenshot of it.
func screenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (error error) {
			*imageBuf, error = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			return error
		}),
	}
}
