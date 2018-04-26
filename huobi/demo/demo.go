package main

import (
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bubble501/sheep/common"
	"github.com/bubble501/sheep/huobi"
	"github.com/bubble501/sheep/okex"
)

func symbolToTopic(first string, second string, market string) string {
	if market == "huobi" {
		return "market." + first + second + ".depth.step0"
	}

	if market == "okex" {
		return "ok_sub_spot_" + first + "_" + second + "_depth_5"
	}

	return ""
}

func topicTosymbol(topic string, market string) string {

	if market == "huobi" {
		items := strings.Split(topic, ".")
		return items[1]
	}
	if market == "okex" {
		items := strings.Split(topic, "_")
		return items[3] + items[4]
	}

	return ""

}

func createListener(market string, manager *common.OrderBookManager) func(topic string, json *simplejson.Json) {
	return func(topic string, json *simplejson.Json) {
		symbol := topicTosymbol(topic, market)
		manager.AddOrderBook(market, symbol, json)
	}
}

func main() {
	markets := []string{"okex", "huobi"}
	//symbols := []string{"btcusdt", "ethusdt"}
	//symbols := []string{"ethusdt", "btcusdt"}
	symbols := []string{"btcusdt", "ethusdt"}
	orderbookManager, _ := common.NewOrderBookManager()
	orderbookManager.InitBook(markets, symbols)
	arbitrator := common.NewArbitrator(orderbookManager)
	orderbookManager.SubscirbeNewOrderEvent(arbitrator)

	ok, err := okex.NewOKEX("", "")
	if err != nil {
		println(err)
	}

	ok.SubscribeDepthDirect(createListener("okex", orderbookManager), symbols)

	hb, err := huobi.NewHuobi("", "")
	if err != nil {
		println(err)
	}

	hb.SubscribeDepthDirect(createListener("huobi", orderbookManager), symbols)

	time.Sleep(time.Hour)
}
