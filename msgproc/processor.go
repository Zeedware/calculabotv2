package msgproc

import "github.com/thoas/go-funk"

var commands []Command = []Command{
	ImageCommand{},
}

func ProcessMessage(update BotUpdate) []BotReply {
	filterFunction := func(command Command) bool {
		return command.IsExecuted(update)
	}
	mapFunction := func(command Command) BotReply {
		return command.Process(update)
	}
	return funk.Map(funk.Filter(commands, filterFunction), mapFunction).([]BotReply)
}
