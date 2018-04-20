package common

import (
	"sync"
	"time"
)

//Order include price and volume.
type Order struct {
	price float32
	volume float32
	key string
}

//Orderbook include bids, asks and timestamps.
type Orderbook struct {
	bids [] Order
	asks [] Order
	ts time.Time
}

//OrderbookMap include the orderbook indexed by exchange and symbol.
type OrderbookMap map[string]map[string]*Orderbook

var instance *OrderbookMap
var once sync.Once

//GetOrderbookMap get the orderbookmap singleton instance.
func GetOrderbookMap() *OrderbookMap {
	once.Do(func() {
		instance = &OrderbookMap{}
		(*instance)["huobi"] = make(map[string]*Orderbook)
		(*instance)["okex"] = make(map[string]*Orderbook)
	})
	return instance
}

//ReplaceOrderbook replace the specified orderbook based on market and symbol. 
func ReplaceOrderbook(market string, symbol string, orderbook *Orderbook) (error) {
	obm := GetOrderbookMap()
	(*obm)[market][symbol] = orderbook
	return nil
}

func merge(first []Order, second []Order) ([]Order, error) {
	firstIndex := 0;
	secondIndex := 0;
	result := make([]Order, 0)
	for firstIndex < len(first) && secondIndex < len(second) {
		if first[firstIndex].key < second[secondIndex].price {

		}
	}

	return nil, nil
}

//MergeOrderbook merge the specified orderbook based on market and symbol.
func MergeOrderbook(market string, symbol string, orderbook *Orderbook) (error) {
	obm := GetOrderbookMap()
	ob := (*obm)[market][symbol]
	(*ob).asks = merge(ob.aks, orderbook.asks)

	
	return nil
}
