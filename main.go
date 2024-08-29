package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func sendMessaggeToLineBot(message string) {
	//channel_secret := os.Getenv("CHANNEL_SECRET")
	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	user_id := os.Getenv("USER_ID")

	bot, err := messaging_api.NewMessagingApiAPI(channel_access_token)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: user_id,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: message,
				},
			},
		},
		"",
	)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Send message to Line Bot!")
	}
}

func requestToWebServer() {
	url := os.Getenv("URL")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Success!")
		sendMessaggeToLineBot("The web server is running!")
	} else if resp.StatusCode == http.StatusForbidden {
		log.Println("Forbidden!")
		sendMessaggeToLineBot("The web server is not running!")
	} else {
		log.Println("Error!")
		sendMessaggeToLineBot("The web server is not running!")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Server started!")

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	requestToWebServer()

	for {
		<-ticker.C
		requestToWebServer()
	}

}
