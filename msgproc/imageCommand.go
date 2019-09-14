package msgproc

import (
	"calculabotv2/qwantservice"
	"regexp"
	"strings"
)

type ImageCommand struct {
}

func (imageCommand ImageCommand) IsExecuted(update BotUpdate) bool {
	regex := regexp.MustCompile("(?i)\\Agbr [[:word:]]+")
	return regex.MatchString(update.Message())
}

func (imageCommand ImageCommand) Process(update BotUpdate) BotReply {
	imageName := getCommandArguments(update.Message())
	_, image, err := qwantservice.GetImage(imageName)
	if err != nil {
		return NewBotMsgReply(update, "image not found")
	}
	return NewBotImgReply(update, image, imageName)
}

func getCommandArguments(input string) string {
	return strings.Split(strings.ToLower(input), "gbr ")[1]
}
