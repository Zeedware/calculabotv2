package msgproc

type Command interface {
	IsExecuted(BotUpdate) bool
	Process(BotUpdate) BotReply
}

type BotUpdate interface {
	Message() string
}

type BotReply interface {
}

type BotMsgReply struct {
}

type BotImgReply struct {
}
