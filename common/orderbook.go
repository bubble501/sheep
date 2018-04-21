package common

import (
	"fmt"
	"time"

	"github.com/bitly/go-simplejson"
)

//Order include price and volume.
type Order struct {
	price  float32
	volume float32
	key    string
}

//OrderBook include bids, asks and timestamps.
type OrderBook struct {
	bids []Order
	asks []Order
	ts   time.Time
}

type JsonAddOrderBook struct {
	market string
	symbol string
	data   *simplejson.Json
}

type JsonAddOrderBookChan = chan JsonAddOrderBook

//OrderBookManager is used to manage all orderbooks.
type OrderBookManager struct {
	books       map[string]map[string]*OrderBook
	jsonAddChan JsonAddOrderBookChan
}

// NewOrderBookManager create OrderBookManager.
func NewOrderBookManager() (manager *OrderBookManager, err error) {
	manager = &OrderBookManager{
		books:       make(map[string]map[string]*OrderBook),
		jsonAddChan: make(JsonAddOrderBookChan, 100),
	}
	manager.books["huobi"] = make(map[string]*OrderBook)
	manager.books["okex"] = make(map[string]*OrderBook)
	go manager.handleJSONAdd()
	return manager, nil
}

//AddOrderBook add orderbook according specified market and symbol.
func (m *OrderBookManager) AddOrderBook(market string, symbol string, data interface{}) error {
	switch t := data.(type) {
	case *simplejson.Json:
		m.jsonAddChan <- JsonAddOrderBook{
			market: market,
			symbol: symbol,
			data:   t,
		}
	}
	return nil
}

func (m *OrderBookManager) handleJSONAdd() {
	for jsonAddOrderBook := range m.jsonAddChan {
		fmt.Printf("add orderbook is here %+v\n", jsonAddOrderBook)
	}
}

// func GetOrderBookManager() *OrderBookManager {
// 	once.Do(func() {
// 		obmInstance = &OrderBookManager{

// 		}

// 	})

// 	return obmInstance;
// }

// //GetOrderbookMap get the orderbookmap singleton instance.
// func GetOrderbookMap() *OrderbookMap {
// 	once.Do(func() {
// 		instance = &OrderbookMap{}
// 		(*instance)["huobi"] = make(map[string]*Orderbook)
// 		(*instance)["okex"] = make(map[string]*Orderbook)
// 	})
// 	return instance
// }

// //ReplaceOrderbook replace the specified orderbook based on market and symbol.
// func ReplaceOrderbook(market string, symbol string, orderbook *Orderbook) (error) {
// 	obm := GetOrderbookMap()
// 	(*obm)[market][symbol] = orderbook
// 	return nil
// }

// func merge(first []Order, second []Order) ([]Order, error) {
// 	firstIndex := 0;
// 	secondIndex := 0;
// 	result := make([]Order, 0)
// 	for firstIndex < len(first) && secondIndex < len(second) {
// 		if first[firstIndex].key < second[secondIndex].price {

// 		}
// 	}

// 	return nil, nil
// }

// //MergeOrderbook merge the specified orderbook based on market and symbol.
// func MergeOrderbook(market string, symbol string, orderbook *Orderbook) (error) {
// 	obm := GetOrderbookMap()
// 	ob := (*obm)[market][symbol]
// 	(*ob).asks = merge(ob.aks, orderbook.asks)

// 	return nil
// }
