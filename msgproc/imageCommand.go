package msgproc

import (
	"calculabotv2/qwantservice"
	"regexp"
	"strings"
)

type ImageCommand struct {
}

func (imageCommand ImageCommand) IsExecuted(update BotUpdate) bool {
	regex := regexp.MustCompile("\\Agbr [[:word:]]+")
	return regex.MatchString(update.Message())
}

func (imageCommand ImageCommand) Process(update BotUpdate) BotReply {
	title, image, err := qwantservice.GetImage(getCommandArguments(update.Message()))
	if err != nil {
		return NewBotMsgReply(update, "image not found")
	}
	return NewBotImgReply(update, image, title)
}

func getCommandArguments(input string) string {
	return strings.Split(input, "gbr ")[1]
}
