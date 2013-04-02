package goweixin

import (
	"crypto/sha1"
	"fmt"
	"github.com/clbanning/x2j"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

var _Debug = false
var DevMode = false

const (
	TEXT     = "text"
	IMAGE    = "image"
	LOCATION = "location"
	LINK     = "link"
	EVENT    = "event"
	VOICE    = "voice"
)

type WxHttpHandler struct {
	Token   string
	Handler WxHandler
}

type WxHandler interface {
	Text(Message) Replay
	Image(Message) Replay
	Location(Message) Replay
	Link(Message) Replay
	Event(Message) Replay
	Voice(Message) Replay
	Default(Message) Replay
}

func (wx *WxHttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var msgType string
	var ok bool
	if _Debug {
		start := time.Now()
		defer func() {
			if ok {
				log.Printf("OK : %s , %ds", msgType, time.Now().Unix()-start.Unix())
			} else {
				log.Printf("ERR: %v", req)
			}
		}()
	}
	if err := req.ParseForm(); err != nil {
		log.Println("Bad Req", req)
		rw.WriteHeader(500)
		return
	}
	if !Verify(wx.Token, req.FormValue("timestamp"), req.FormValue("nonce"), req.FormValue("signature")) {
		rw.WriteHeader(403)
		return
	}
	if req.Method == "GET" {
		rw.Write([]byte(req.FormValue("echostr")))
		return
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Req Body Read ERR", req, err)
		rw.WriteHeader(500)
		return
	}
	if DevMode {
		log.Println("Req\n" + string(data))
	}

	root, err := x2j.DocToMap(string(data))
	if err != nil {
		fmt.Println("Bad XML Req", err)
		return
	}
	msg := Message(root["xml"].(map[string]interface{}))
	msgType = msg.MsgType()
	if _Debug {
		log.Println("MsgType =", msgType)
	}
	var reply Replay
	switch msgType {
	case TEXT:
		reply = wx.Handler.Text(msg)
	case IMAGE:
		reply = wx.Handler.Image(msg)
	case LOCATION:
		reply = wx.Handler.Location(msg)
	case LINK:
		reply = wx.Handler.Link(msg)
	case EVENT:
		reply = wx.Handler.Event(msg)
	case VOICE:
		reply = wx.Handler.Voice(msg)
	default:
		reply = wx.Handler.Default(msg)
	}
	if reply == nil {
		ok = true
		if _Debug {
			log.Println("Reply nil")
		}
		return // http 200
	}

	// auto-fix
	if reply.FromUserName() == "" {
		reply.SetFromUserName(msg.ToUserName())
	}
	if reply.ToUserName() == "" {
		reply.SetToUserName(msg.FromUserName())
	}
	if reply.MsgType() == "" {
		reply.SetMsgType(TEXT)
	}
	reply.SetCreateTime(time.Now().Unix())

	if _, ok = reply["FuncFlag"]; !ok {
		reply.SetFuncFlag(0)
	}

	rw.Write([]byte("<xml>"))
	_re := MapToXmlString(reply)
	if _Debug {
		log.Println("Reply\n" + _re)
	}
	rw.Write([]byte(_re))
	rw.Write([]byte("</xml>"))
	ok = true
}

func Verify(token string, timestamp string, nonce string, signature string) bool {
	if DevMode {
		return true
	}
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)
	key := strs[0] + strs[1] + strs[2]
	if _Debug {
		log.Println("Verify key=", key)
	}
	h := sha1.New()
	h.Write([]byte(key))
	re := fmt.Sprintf("%x", h.Sum(nil))
	if _Debug {
		log.Println("Verify", signature, re)
	}
	return signature == re
}

type BaseWeiXinHandler struct {
}

func (h *BaseWeiXinHandler) Text(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Image(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Location(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Link(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Event(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Voice(msg Message) Replay {
	return h.Default(msg)
}
func (h *BaseWeiXinHandler) Default(msg Message) Replay {
	return nil
}

func SetDebug(_debug bool) {
	_Debug = _debug
}

type XmlMessage struct {
	Msg Message `xml:"xml"`
}
