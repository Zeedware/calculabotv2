package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	TelegramBot *tgbotapi.BotAPI
}

func NewBot(telegramToken string) *Bot {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	return &Bot{
		TelegramBot: bot,
	}
}

func (bot *Bot) Start() {

	tBot := bot.TelegramBot
	log.Printf("Authorized on account %s", tBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := tBot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		_ = transformTelegramUpdateToBotUpdate(&update)

	}
}
