package main

import (
	"fmt"
	"time"

	"github.com/bubble501/sheep/huobi"
)

func main() {

	// orderbookManager, _ := common.NewOrderBookManager()

	// topicToSymbol := map[string]string{
	// 	"ok_sub_spot_btc_usdt_depth_5": "btcusdt",
	// }
	// fmt.Println("shit")
	// okex, err := okex.NewMarket()
	// if err != nil {
	// 	println(err)
	// }

	// spotOrderbookdepthListener := func(topic string, json *simplejson.Json) {
	// 	symbol := topicToSymbol[topic]
	// 	orderbookManager.AddOrderBook("okex", symbol, json)
	// }

	// okex.Subscribe("ok_sub_spot_btc_usdt_depth_5", spotOrderbookdepthListener)

	// h, err := huobi.NewHuobi("", "")
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// // 打开websocket通信
	// err = h.OpenWebsocket()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer h.CloseWebsocket()

	// //获取账户余额
	// balances, err := h.GetAccountBalance()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(balances)

	//webcosket监听函数
	listen := func(symbol string, depth *huobi.MarketDepth) {
		//orderbookManager.AddOrderBook("huobi", symbol, depth)
		fmt.Println("shit")
	}

	//设置监听
	h.SetDepthlListener(listen)

	//订阅
	h.SubscribeDepth("btcusdt")

	time.Sleep(time.Hour)
}
