package msgproc

import (
	"regexp"
)

type ImageCommand struct {
}

func (imageCommand ImageCommand) IsExecuted(update BotUpdate) bool {
	regex := regexp.MustCompile("\\Agbr [[:word:]]+")
	return regex.MatchString(update.Message())
}

func (imageCommand ImageCommand) Process(update BotUpdate) BotReply {
	return nil
}
