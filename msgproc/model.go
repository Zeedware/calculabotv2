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
	msg string
}

func NewBotMsgReply(msg string) BotMsgReply {
	return BotMsgReply{msg: msg}
}

type BotImgReply struct {
	image []byte
}

func NewBotImgReply(image []byte) BotImgReply {
	return BotImgReply{image: image}
}
