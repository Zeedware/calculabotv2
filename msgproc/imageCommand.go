package msgproc

import (
	"calculabotv2/qwantservice"
	"regexp"
)

type ImageCommand struct {
}

func (imageCommand ImageCommand) IsExecuted(update BotUpdate) bool {
	regex := regexp.MustCompile("\\Agbr [[:word:]]+")
	return regex.MatchString(update.Message())
}

func (imageCommand ImageCommand) Process(update BotUpdate) BotReply {
	_, image, err := qwantservice.GetImage(update.Message())
	if err != nil {
		return NewBotMsgReply("image not found")
	}
	return NewBotImgReply(image)
}
