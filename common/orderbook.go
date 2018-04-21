package common

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bitly/go-simplejson"
)

//Order include price and volume.
type Order struct {
	price  string
	volume string
}

//OrderBook include bids, asks and timestamps.
type OrderBook struct {
	bids []Order
	asks []Order
	ts   time.Time
}

//JSONAddOrderBook is an object used to define JsonAddOrderBookChan.
type JSONAddOrderBook struct {
	market string
	symbol string
	data   *simplejson.Json
}

type JsonAddOrderBookChan = chan JSONAddOrderBook

//OrderBookManager is used to manage all orderbooks.
type OrderBookManager struct {
	mutex *sync.RWMutex
	//the key is the join of market and symbol with ":"
	books       map[string]*OrderBook
	jsonAddChan JsonAddOrderBookChan
}

// NewOrderBookManager create OrderBookManager.
func NewOrderBookManager() (manager *OrderBookManager, err error) {
	manager = &OrderBookManager{
		books:       make(map[string]*OrderBook, MarketSize*SymbolSize),
		jsonAddChan: make(JsonAddOrderBookChan, 100),
		mutex:       &sync.RWMutex{},
	}

	manager.books["huobi:btcusdt"] = &OrderBook{
		bids: make([]Order, OrderBookDepth),
		asks: make([]Order, OrderBookDepth),
	}

	manager.books["okex:btcusdt"] = &OrderBook{
		bids: make([]Order, OrderBookDepth),
		asks: make([]Order, OrderBookDepth),
	}
	go manager.handleJSONAdd()
	return manager, nil
}

//AddOrderBook add orderbook according specified market and symbol.
func (m *OrderBookManager) AddOrderBook(market string, symbol string, data interface{}) error {
	switch t := data.(type) {
	case *simplejson.Json:
		m.jsonAddChan <- JSONAddOrderBook{
			market: market,
			symbol: symbol,
			data:   t,
		}
	}
	return nil
}

func (m *OrderBookManager) handleJSONAdd() {
	for chanItem := range m.jsonAddChan {
		s := []string{chanItem.market, chanItem.symbol}
		key := strings.Join(s, ":")

		fmt.Printf("add orderbook is here %+v\n", chanItem)
		data := chanItem.data.GetIndex(0).Get("data")
		timestamp := chanItem.data.GetIndex(0).Get("timestamp").MustUint64()
		asks := data.Get("asks")
		length := len(asks.MustArray())
		fmt.Printf("the length is %d", length)
		m.mutex.Lock()
		defer m.mutex.Unlock()
		for i := 0; i < length; i++ {
			pairs := asks.GetIndex(i).MustStringArray()
			m.books[key].asks[i].price = pairs[0]
			m.books[key].asks[i].volume = pairs[1]
		}

		bids := data.Get("bids")
		length = len(bids.MustArray())

		for i := 0; i < length; i++ {
			pairs := bids.GetIndex(i).MustStringArray()
			m.books[key].bids[i].price = pairs[0]
			m.books[key].bids[i].volume = pairs[1]
		}
		m.books[key].ts = time.Unix(int64(timestamp), 0)
		fmt.Printf("#########%+v\n", *m)
	}
}

// func (m *OrderBookManager) String() string {
// 	m.mutex.RLock()
// 	defer m.mutex.RUnlock()

// }
