package qwantservice

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	title, image, err := GetImage("hell")
	Nil(t, err)
	NotEmpty(t, title)
	NotEmpty(t, image)
}
