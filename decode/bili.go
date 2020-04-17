package decode

import (
	"encoding/json"
	"errors"
	"strings"
)

//ErrOtherMsg 其他消息，指非rich类型
var ErrOtherMsg = errors.New("The msg isn't a rich")

//ErrOtherCard 其他卡片，指非Bilibili卡片
var ErrOtherCard = errors.New("The card isn't send by bilibili")

//BiliCard bilibili小程序卡片
type BiliCard struct {
	Title   string
	Content bilicontent
}

type bilicontent struct {
	Detail1 struct {
		Appid string `json:"appid"`
		Desc  string `json:"desc"`
		Host  struct {
			Nick string `json:"nick"`
			Uin  int64  `json:"uin"`
		} `json:"host"`
		Icon               string      `json:"icon"`
		Rreview            string      `json:"preview"`
		QQdocurl           string      `json:"qqdocurl"`
		Ccene              string      `json:"scene"`
		ShareTemlplateData interface{} `json:"shareTemlplateData"`
		ShareTemlplateID   string      `json:"shareTemlplateId"`
		Title              string      `json:"title"`
		URL                string      `json:"url"`
	} `json:"detail_1"`
}

//Bili 解析bilibili卡片消息
func Bili(msg string) (BiliCard, error) {
	var card BiliCard
	if len(msg) < 26 {
		return card, ErrOtherMsg
	}
	if msg[:9] != "[CQ:rich," {
		return card, ErrOtherMsg
	}
	//len("[CQ:rich,title=")==15
	l := 15
	r := 15
	for r < len(msg)-8 {
		if msg[r] == ',' {
			break
		}
		r++
	}
	card.Title = cqUnEscaper.Replace(msg[l:r])
	if card.Title != "[QQ小程序]哔哩哔哩" {
		return card, ErrOtherCard
	}
	bilijson := cqUnEscaper.Replace(msg[r+9 : len(msg)-1])
	json.Unmarshal([]byte(bilijson), &(card.Content))
	return card, nil
}

var cqUnEscaper = strings.NewReplacer(
	"&amp;", "&",
	"&#91;", "[",
	"&#93;", "]",
	"&#44;", ",",
)
