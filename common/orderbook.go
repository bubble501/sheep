package common

import (
	"fmt"
	"strconv"
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
		bids := data.Get("bids")
		length := len(asks.MustArray())
		m.updateOrderBook(key, asks, bids, int64(timestamp), length)
	}
}

func (m *OrderBookManager) updateOrderBook(key string, asks *simplejson.Json, bids *simplejson.Json, timestamp int64, length int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i := 0; i < length; i++ {
		ask := asks.GetIndex(i).MustStringArray()
		bid := bids.GetIndex(i).MustStringArray()
		m.books[key].asks[length-i-1].price = ask[0]
		m.books[key].asks[length-i-1].volume = ask[1]
		m.books[key].bids[i].price = bid[0]
		m.books[key].bids[i].volume = bid[1]

	}
	m.books[key].ts = time.Unix(timestamp, 0)
	fmt.Println(m)
	return nil
}

func (m *OrderBookManager) String() string {
	result := make([]string, 0)
	for k, v := range m.books {
		result = append(result, k)
		for i, ask := range v.asks {
			line := "ask" + strconv.Itoa(i) + "[" + ask.price + "," + ask.volume + "]"
			result = append(result, line)
		}

		for i, bid := range v.bids {
			line := "bid" + strconv.Itoa(i) + "[" + bid.price + "," + bid.volume + "]"
			result = append(result, line)
		}

	}
	return strings.Join(result, "\n")
}
