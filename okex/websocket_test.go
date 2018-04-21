package okex

import (
	"fmt"
	"testing"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bubble501/sheep/common"
	"github.com/stretchr/testify/assert"
)

func TestOrderBookGetInstance(t *testing.T) {
	topicToSymbol := map[string]string{
		"ok_sub_spot_btc_usdt_depth_5": "btcusdt",
	}
	orderbookManager, _ := common.NewOrderBookManager()
	fmt.Println("shit")
	h, err := NewMarket()
	if err != nil {
		println(err)
	}
	assert.Equal(t, err, nil)
	spotOrderbookdepthListener := func(topic string, json *simplejson.Json) {
		symbol := topicToSymbol[topic]
		orderbookManager.AddOrderBook("okex", symbol, json)
		println("******** processed ****************")
	}

	h.Subscribe("ok_sub_spot_btc_usdt_depth_5", spotOrderbookdepthListener)
	time.Sleep(time.Hour)
}
