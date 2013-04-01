package goweixin

type TxtIn struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

type PicIn struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	PicUrl       string
	MsgId        int64
}

type LocIn struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Location_X   string
	Location_Y   string
	Scale        int
	Label        string
	MsgId        int64
}

type PushIn struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
}

//------------------------------
type TxtOut struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	FuncFlag     int
}

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
