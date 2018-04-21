package main

import (
	"log"
	"time"

	"github.com/bubble501/sheep/huobi"
)

func main() {
	h, err := huobi.NewHuobi("", "")
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 打开websocket通信
	err = h.OpenWebsocket()
	if err != nil {
		log.Fatal(err)
	}
	defer h.CloseWebsocket()

	// //获取账户余额
	// balances, err := h.GetAccountBalance()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(balances)

	//webcosket监听函数
	listen := func(symbol string, depth *huobi.MarketDepth) {
		log.Println(depth)
	}

	//设置监听
	h.SetDepthlListener(listen)

	//订阅
	h.SubscribeDepth("btcusdt")
	h.SubscribeDepth("ethusdt")

	time.Sleep(time.Hour)
}
