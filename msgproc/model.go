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
	botUpdate BotUpdate
	msg       string
}

func NewBotMsgReply(botUpdate BotUpdate, msg string) BotMsgReply {
	return BotMsgReply{botUpdate: botUpdate, msg: msg}
}

type BotImgReply struct {
	botUpdate BotUpdate
	image     []byte
	caption   string
}

func NewBotImgReply(botUpdate BotUpdate, image []byte, caption string) BotImgReply {
	return BotImgReply{botUpdate: botUpdate, image: image, caption: caption}
}
