package goweixin

import (
	"strconv"
)

type Message map[string]interface{}

//------------------------------------
func (w Message) String(key string) string {
	if str, ok := w[key]; ok {
		return str.(string)
	}
	return ""
}

func (w Message) Int64(key string) int64 {
	if val, ok := w[key]; ok {
		switch val.(type) {
		case string:
			i, _ := strconv.ParseInt(val.(string), 0, 64)
			return i
		case int:
			return int64(val.(int))
		case int64:
			return val.(int64)
		}
	}
	return 0
}

//------------------------------------

func (w Message) ToUserName() string {
	return w.String("ToUserName")
}
func (w Message) FromUserName() string {
	return w.String("FromUserName")
}
func (w Message) CreateTime() int64 {
	return w.Int64("CreateTime")
}
func (w Message) MsgType() string {
	return w.String("MsgType")
}
func (w Message) MsgId() string {
	return w.String("MsgId")
}

//------------------------------------

func (w Message) Content() string {
	return w.String("Content")
}

//-----------------------------------

func (w Message) PicUrl() string {
	return w.String("PicUrl")
}

//-----------------------------------

func (w Message) Location_X() string {
	return w.String("Location_X")
}
func (w Message) Location_Y() string {
	return w.String("Location_Y")
}
func (w Message) Scale() int64 {
	return w.Int64("Scale")
}
func (w Message) Label() string {
	return w.String("Label")
}

//--------------------------------

func (w Message) Event() int64 {
	return w.Int64("Event")
}
func (w Message) EventKey() string {
	return w.String("EventKey")
}

//--------------------------------
func (w Message) Title() int64 {
	return w.Int64("Title")
}
func (w Message) Description() string {
	return w.String("Description")
}
func (w Message) Url() string {
	return w.String("Url")
}

//------------------------------
//------------------------------
//------------------------------

type Replay map[string]interface{}

func (r Replay) String(key string) string {
	if str, ok := r[key]; ok {
		return str.(string)
	}
	return ""
}

func (r Replay) Int64(key string) int64 {
	if val, ok := r[key]; ok {
		switch val.(type) {
		case string:
			i, _ := strconv.ParseInt(val.(string), 0, 64)
			return i
		case int:
			return int64(val.(int))
		case int64:
			return val.(int64)
		}
	}
	return 0
}

func (r Replay) ToUserName() string {
	return r.String("ToUserName")
}
func (r Replay) FromUserName() string {
	return r.String("FromUserName")
}
func (r Replay) CreateTime() int64 {
	return r.Int64("CreateTime")
}
func (r Replay) MsgType() string {
	return r.String("MsgType")
}
func (r Replay) FuncFlag() int64 {
	return r.Int64("FuncFlag")
}

func (r Replay) SetToUserName(val string) Replay {
	r["ToUserName"] = val
	return r
}
func (r Replay) SetFromUserName(val string) Replay {
	r["FromUserName"] = val
	return r
}
func (r Replay) SetCreateTime(val int64) Replay {
	r["CreateTime"] = val
	return r
}
func (r Replay) SetMsgType(val string) Replay {
	r["MsgType"] = val
	return r
}
func (r Replay) SetFuncFlag(val int64) Replay {
	r["FuncFlag"] = val
	return r
}

//----------------------------------------
func (r Replay) Content() string {
	return r.String("Content")
}
func (r Replay) SetContent(val string) Replay {
	r["Content"] = val
	return r
}

//----------------------------------------

type MusicOut struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string

	Title       string `xml:"Music>Title"`
	Description string `xml:"Music>Description"`
	MusicUrl    string `xml:"Music>MusicUrl"`
	HQMusicUrl  string `xml:"Music>HQMusicUrl"`
	FuncFlag    int
}
