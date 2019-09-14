package msgproc

type Command interface {
	IsExecuted(BotUpdate) bool
	Process(BotUpdate) BotReply
}

type BotUpdate interface {
	Message() string
}

type BotReply interface {
	BotUpdate() BotUpdate
}

type BotMsgReply struct {
	botUpdate BotUpdate
	message   string
}

func (b BotMsgReply) Message() string {
	return b.message
}

func (b BotMsgReply) BotUpdate() BotUpdate {
	return b.botUpdate
}

func NewBotMsgReply(botUpdate BotUpdate, message string) BotMsgReply {
	return BotMsgReply{botUpdate: botUpdate, message: message}
}

type BotImgReply struct {
	botUpdate BotUpdate
	image     []byte
	caption   string
}

func (b BotImgReply) Caption() string {
	return b.caption
}

func (b BotImgReply) Image() []byte {
	return b.image
}

func (b BotImgReply) BotUpdate() BotUpdate {
	return b.botUpdate
}

func NewBotImgReply(botUpdate BotUpdate, image []byte, caption string) BotImgReply {
	return BotImgReply{botUpdate: botUpdate, image: image, caption: caption}
}
