package main

import (
	"fmt"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bubble501/sheep/common"
	"github.com/bubble501/sheep/huobi"
)

func main() {

	orderbookManager, _ := common.NewOrderBookManager()

	topicToSymbol := map[string]string{
		"market.btcusdt.depth.step0": "btcusdt",
	}
	fmt.Println("shit")
	market, err := huobi.NewMarket()
	if err != nil {
		println(err)
	}

	spotOrderbookdepthListener := func(topic string, json *simplejson.Json) {
		symbol := topicToSymbol[topic]
		orderbookManager.AddOrderBook("huobi", symbol, json)
	}

	market.Subscribe("market.btcusdt.depth.step0", spotOrderbookdepthListener)

	time.Sleep(time.Hour)
}
