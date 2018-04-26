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
	bids   []Order
	asks   []Order
	ts     time.Time
	length int
}

//JSONAddOrderBook is an object used to define JsonAddOrderBookChan.
type JSONAddOrderBook struct {
	market string
	symbol string
	data   *simplejson.Json
}

type JsonAddOrderBookChan = chan JSONAddOrderBook

type NewOrderEventHandler interface {
	handleNewOrder(oderbook *JSONAddOrderBook)
}

//OrderBookManager is used to manage all orderbooks.
type OrderBookManager struct {
	mutex *sync.RWMutex
	//the key is the join of market and symbol with ":"
	books       map[string]*OrderBook
	jsonAddChan JsonAddOrderBookChan
	noeHandler  NewOrderEventHandler
	markets     []string
	symbols     []string
}

// NewOrderBookManager create OrderBookManager.
func NewOrderBookManager() (manager *OrderBookManager, err error) {
	manager = &OrderBookManager{
		books:       make(map[string]*OrderBook, MarketSize*SymbolSize),
		jsonAddChan: make(JsonAddOrderBookChan, 100),
		mutex:       &sync.RWMutex{},
	}
	go manager.handleJSONAdd()
	return manager, nil
}

func (m *OrderBookManager) SubscirbeNewOrderEvent(handler NewOrderEventHandler) {
	m.noeHandler = handler
}

func (m *OrderBookManager) InitBook(markets []string, pairs []string) {
	m.markets = make([]string, len(markets))
	m.symbols = make([]string, len(pairs))
	copy(m.markets, markets)
	copy(m.symbols, pairs)
	for _, market := range markets {
		for _, pair := range pairs {
			key := market + ":" + pair
			m.books[key] = &OrderBook{
				bids: make([]Order, OrderBookDepth),
				asks: make([]Order, OrderBookDepth),
			}
		}
	}
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
		switch chanItem.market {
		case "okex":
			m.handleOkexJSONAdd(&chanItem)
		case "huobi":
			m.handleHuobiJSONAdd(&chanItem)
		}
		m.noeHandler.handleNewOrder(&chanItem)
	}
}

func (m *OrderBookManager) handleHuobiJSONAdd(chanItem *JSONAddOrderBook) {
	fmt.Println(chanItem.market)
	fmt.Println(chanItem.symbol)
	s := []string{chanItem.market, chanItem.symbol}
	key := strings.Join(s, ":")
	fmt.Println(key)
	tick := chanItem.data.Get("tick")
	timestamp := chanItem.data.Get("ts").MustInt64()
	asks := tick.Get("asks")
	bids := tick.Get("bids")
	length := len(asks.MustArray())
	if length > OrderBookDepth {
		length = OrderBookDepth
	}

	m.updateHuobiOrderBook(key, asks, bids, timestamp, length)
}

func (m *OrderBookManager) updateHuobiOrderBook(key string, asks *simplejson.Json, bids *simplejson.Json, timestamp int64, length int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i := 0; i < length; i++ {
		askPrice := asks.GetIndex(i).GetIndex(0).MustFloat64()
		askVolume := asks.GetIndex(i).GetIndex(1).MustFloat64()
		askPriceStr := strconv.FormatFloat(askPrice, 'f', 6, 64)
		askVolumeStr := strconv.FormatFloat(askVolume, 'f', 6, 64)
		m.books[key].asks[i].price = askPriceStr
		m.books[key].asks[i].volume = askVolumeStr

		bidPrice := bids.GetIndex(i).GetIndex(0).MustFloat64()
		bidVolume := bids.GetIndex(i).GetIndex(1).MustFloat64()
		bidPriceStr := strconv.FormatFloat(bidPrice, 'f', 6, 64)
		bidVolumeStr := strconv.FormatFloat(bidVolume, 'f', 6, 64)
		m.books[key].bids[i].price = bidPriceStr
		m.books[key].bids[i].volume = bidVolumeStr
	}
	m.books[key].length = length
	m.books[key].ts = time.Unix(timestamp, 0)
	//fmt.Println(m)
	return nil
}

func (m *OrderBookManager) handleOkexJSONAdd(chanItem *JSONAddOrderBook) {
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
	m.books[key].length = length
	//fmt.Println(m)
	return nil
}

func (m *OrderBookManager) String() string {
	result := make([]string, 0)
	for k, v := range m.books {
		result = append(result, k)
		result = append(result, v.ts.String())
		for i := 0; i < v.length; i++ {
			ask := v.asks[i]
			line := "ask" + strconv.Itoa(i) + "[" + ask.price + "," + ask.volume + "]"
			result = append(result, line)
		}
		for i := 0; i < v.length; i++ {
			bid := v.bids[i]
			line := "bid" + strconv.Itoa(i) + "[" + bid.price + "," + bid.volume + "]"
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}
