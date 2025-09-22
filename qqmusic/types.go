package qqmusic

import "errors"

var (
	ErrSearchSong      = errors.New("[QQMusicAdapter] Failed to search song")
	ErrGetPlayURL      = errors.New("[QQMusicAdapter] Failed to get play url")
	ErrEmptyPlayURL    = errors.New("[QQMusicAdapter] play url is empty")
	ErrGetAlbumDetail  = errors.New("[QQMusicAdapter] Failed to album detail")
	ErrParseTimeString = errors.New("[QQMusicAdapter] Failed to parse time string")
	ErrStatusCode      = errors.New("[QQMusicAdapter] Http status error")
	ErrUnmarshaling    = errors.New("[QQMusicAdapter] Failed to unmarshaling body bytes")
	ErrBase64Decode    = errors.New("[QQMusicAdapter] Failed to decode base64")
)

// 搜索结果响应体结构体
type searchResponseBody struct {
	Code int64 `json:"code"`
	Data struct {
		Keyword string `json:"keyword"`
		Song    struct {
			List []songItem `json:"list"`
		} `json:"song"`
	} `json:"data"`
}

type songItem struct {
	Albumid   int64        `json:"albumid"`
	Albummid  string       `json:"albummid"`
	Albumname string       `json:"albumname"`
	Docid     string       `json:"docid"`
	Pubtime   int64        `json:"pubtime"`
	Size128   int64        `json:"size128"`
	Size320   int64        `json:"size320"`
	Sizeape   int64        `json:"sizeape"`
	Sizeflac  int64        `json:"sizeflac"`
	Sizeogg   int64        `json:"sizeogg"`
	Songid    int64        `json:"songid"`
	Songmid   string       `json:"songmid"`
	Songname  string       `json:"songname"`
	Stream    int64        `json:"stream"`
	Singer    []singerItem `json:"singer"`
	Pay       struct {
		Payalbum      int64 `json:"payalbum"`
		Payalbumprice int64 `json:"payalbumprice"`
		Paydownload   int64 `json:"paydownload"`
		Payinfo       int64 `json:"payinfo"`
		Payplay       int64 `json:"payplay"`
		Paytrackmouth int64 `json:"paytrackmouth"`
		Paytrackprice int64 `json:"paytrackprice"`
	} `json:"pay"`
}

type singerItem struct {
	Id   int64  `json:"id"`
	Mid  string `json:"mid"`
	Name string `json:"name"`
}

type requestResult struct {
	Code int `json:"code"`
	Req0 struct {
		Code int `json:"code"`
		Data struct {
			MidURLInfo []midURLInfo `json:"midurlinfo"`
			Expiration int          `json:"expiration"`
		} `json:"data"`
	} `json:"req_0"`
}

type midURLInfo struct {
	Songmid  string `json:"songmid"`
	Filename string `json:"filename"`
	Purl     string `json:"purl"`
	Vkey     string `json:"vkey"`
}

type picUrlItem struct {
	PicUrl string `json:"picurl"`
}

type singer struct {
	Singerid   string `json:"singerid"`
	Singermid  string `json:"singermid"`
	Singername string `json:"singername"`
}

type albumDetailResult struct {
	Code int `json:"code"`
	Data struct {
		AlbumMid    string       `json:"album_mid"`
		AlbumName   string       `json:"album_name"`
		Desc        string       `json:"desc"`
		HeadpicList []picUrlItem `json:"headpiclist"`
		PublicTime  string       `json:"publictime"`
		SingerInfo  []singer     `json:"singerinfo"`
		SongList    []songItem   `json:"songlist"`
	} `json:"data"`
}
