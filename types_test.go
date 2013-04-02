package goweixin

import (
	"fmt"
	"github.com/clbanning/x2j"
	"testing"
)

func TestFromXml(*testing.T) {
	str := `
<xml><ToUserName><![CDATA[gh_2dc74cccf555]]></ToUserName>
<FromUserName><![CDATA[oSmHgjkiNii6XnhVXVN5Rj5DDARE]]></FromUserName>
<CreateTime>1364877454</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[  uhhgggg]]></Content>
<MsgId>5862104027977744402</MsgId>
</xml>
	`
	root, err := x2j.DocToMap(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Message(root["xml"].(map[string]interface{})).MsgType())
	fmt.Println(Message(root["xml"].(map[string]interface{})).CreateTime())
}

func TestToXml(*testing.T) {
	str := `
<xml><ToUserName><![CDATA[gh_2dc74cccf555]]></ToUserName>
<FromUserName><![CDATA[oSmHgjkiNii6XnhVXVN5Rj5DDARE]]></FromUserName>
<CreateTime>1364877454</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[  uhhgggg]]></Content>
<MsgId>5862104027977744402</MsgId>
</xml>
	`
	root, err := x2j.DocToMap(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := Message(root["xml"].(map[string]interface{}))
	fmt.Println(msg.MsgType())

	fmt.Println(MapToXmlString(msg))
}
