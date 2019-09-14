package main

import (
	"calculabotv2/msgproc"
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
	u.Timeout = 1
	updates, err := tBot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		go bot.processMessage(update)
	}
}

func (bot *Bot) processMessage(update tgbotapi.Update) {
	botUpdate := transformTelegramUpdateToBotUpdate(update)
	reply := msgproc.ProcessMessage(botUpdate)
	bot.iterateReply(reply)
}

func (bot *Bot) iterateReply(replies []msgproc.BotReply) {
	for _, reply := range replies {
		bot.sendTelegram(reply)
	}
}

func (bot *Bot) sendTelegram(reply msgproc.BotReply) {
	tgUpdate := reply.BotUpdate().(TelegramUpdate)
	var msg tgbotapi.Chattable
	switch reply.(type) {
	case msgproc.BotMsgReply:
		msg = bot.sendMsgReply(reply.(msgproc.BotMsgReply), tgUpdate)
	case msgproc.BotImgReply:
		msg = bot.sendImgReply(reply.(msgproc.BotImgReply), tgUpdate)
	}
	_, err := bot.TelegramBot.Send(msg)
	log.Println(err)
}

func (bot *Bot) sendImgReply(reply msgproc.BotImgReply, tgUpdate TelegramUpdate) tgbotapi.Chattable {
	imgByte := tgbotapi.FileBytes{Name: "image.jpg", Bytes: reply.Image()}
	msg := tgbotapi.NewPhotoUpload(tgUpdate.update.Message.Chat.ID, imgByte)
	msg.Caption = reply.Caption()
	return msg
}

func (bot *Bot) sendMsgReply(reply msgproc.BotMsgReply, tgUpdate TelegramUpdate) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(tgUpdate.update.Message.Chat.ID, reply.Message())
	return msg
}

func (bot *Bot) generateDefaultMsg(tgUpdate TelegramUpdate) tgbotapi.Chattable {
	return tgbotapi.NewMessage(tgUpdate.update.Message.Chat.ID, "Something happened")
}
