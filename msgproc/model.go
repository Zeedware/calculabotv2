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
	update BotUpdate
	msg    string
}

func NewBotMsgReply(update BotUpdate, msg string) BotMsgReply {
	return BotMsgReply{update: update, msg: msg}
}

type BotImgReply struct {
	update  BotUpdate
	image   []byte
	caption string
}

func NewBotImgReply(update BotUpdate, image []byte, caption string) BotImgReply {
	return BotImgReply{update: update, image: image, caption: caption}
}
