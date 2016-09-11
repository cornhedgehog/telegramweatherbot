package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("secret telegram token, no peeping around, please")
	if err != nil {
		log.Panic(err)
	}

	cityCode := "105" //по умолчанию - Калининград

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message.Text == "/start" {
			welcomeMsg := tgbotapi.NewMessage(update.Message.Chat.ID, GetStartMessage(cityCode))
			bot.Send(welcomeMsg)
		}

		if update.Message.Text == "/help" {
			weatherDataMsg := tgbotapi.NewMessage(update.Message.Chat.ID, GetWeatherData(cityCode))
			bot.Send(weatherDataMsg)
		}
	}
}
