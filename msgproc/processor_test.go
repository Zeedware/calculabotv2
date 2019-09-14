package msgproc

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessMessage(t *testing.T) {
	update := TestUpdate{message: "gbr test"}
	botReplies := ProcessMessage(update)
	Len(t, botReplies, 1)
	imgReply, ok := botReplies[0].(BotImgReply)
	True(t, ok)
	NotEmpty(t, imgReply.image)
	NotEmpty(t, imgReply.caption)
}
