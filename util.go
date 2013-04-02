package goweixin

import (
	"fmt"
)

func ReplyText(txt string) Replay {
	replay := Replay{}
	replay.SetContent(txt)
	return replay
}

func ReplyTextf(format string, args ...interface{}) Replay {
	return ReplyText(fmt.Sprintf(format, args...))
}
