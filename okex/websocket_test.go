package okex

import (
	"testing"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
)

func TestOrderBookGetInstance(t *testing.T) {
	topicToSymbol := map[string]string{
		"ok_sub_spot_bch_btc_depth_5": "BCHBTC",
	}

	h, err := NewMarket()
	assert.Equal(t, err, nil)
	spotOrderbookdepthListener := func(topic string, json *simplejson.Json) {
		market := "okex"
		symbol := topicToSymbol[topic]

		println("******** processed ****************")
	}

	h.Subscribe("ok_sub_spot_bch_btc_depth_5", listen)
	time.Sleep(time.Hour)
}
