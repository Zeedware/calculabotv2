package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type BotUpdate interface {
}

type TelegramUpdate struct {
	update *tgbotapi.Update
}

func NewTelegramUpdate(update *tgbotapi.Update) TelegramUpdate {
	return TelegramUpdate{update: update}
}

func transformTelegramUpdateToBotUpdate(update *tgbotapi.Update) TelegramUpdate {
	return NewTelegramUpdate(update)
}
