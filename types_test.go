package goweixin

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestToXml(*testing.T) {
	reply := MusicOut{}
	reply.MsgType = "music"
	reply.HQMusicUrl = "http://wendal.net/XXX"
	data, err := xml.Marshal(reply)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	m := map[string]MusicOut{}
	err = xml.Unmarshal([]byte("<xml>"+string(data)+"</xml>"), &m)
	fmt.Println(err)
	fmt.Println(m)
}
