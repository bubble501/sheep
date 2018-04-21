package main

import (
	"fmt"
	"log"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bubble501/sheep/okex"
)

func main() {
	h, err := okex.NewMarket()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// balances, err := h.GetAccountBalance()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(balances)

	//webcosket监听函数
	listen := func(symbol string, json *simplejson.Json) {
		fmt.Println("shit")
	}

	fmt.Println("before")
	h.Subscribe("ok_sub_spot_bch_btc_depth_5", listen)
	fmt.Println("after")
	time.Sleep(time.Hour)
}
