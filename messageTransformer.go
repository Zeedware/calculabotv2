package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func NewTelegramUpdate(update tgbotapi.Update) TelegramUpdate {
	return TelegramUpdate{update: update}
}

func transformTelegramUpdateToBotUpdate(update tgbotapi.Update) TelegramUpdate {
	return NewTelegramUpdate(update)
}

type TelegramUpdate struct {
	update tgbotapi.Update
}

func (telegramUpdate TelegramUpdate) Message() string {
	return telegramUpdate.update.Message.Text
}

func (telegramUpdate TelegramUpdate) isGroup() bool {
	return true
}
