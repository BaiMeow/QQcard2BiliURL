package main

import (
	"github.com/MscBaiMeow/QQcard2BiliURL/decode"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: BiliURL
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: 将qq卡片的bilibili小程序分享转为bilibili的URL直接发出来
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.biliurl" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if card, err := decode.Bili(msg); err == nil {
		cqp.SendGroupMsg(fromGroup, card.Content.Detail1.QQdocurl)
	}
	return 0
}
