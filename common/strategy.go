package common

import "fmt"
import "strconv"

//Arbitrator is a strategy which move the coin from exchange which has high price
//to the one which has low price.
type Arbitrator struct {
	// accountManager   *AccountManager
	orderBookManager *OrderBookManager

	// orderbookUpdateEvents chan *JSONAddOrderBook
}

func NewArbitrator(obm *OrderBookManager) *Arbitrator {
	return &Arbitrator{
		orderBookManager: obm,
		// orderbookUpdateEvents: make(chan *JSONAddOrderBook, 100),
	}
}

func (m *Arbitrator) handleNewOrder(event *JSONAddOrderBook) {
	fmt.Println("coming from arbitrator..")
	fmt.Println(event.market, event.symbol)
	pair := event.symbol
	newMarket := event.market
	newKey := newMarket + ":" + pair
	newBestBidPrice, _ := strconv.ParseFloat(m.orderBookManager.books[newKey].bids[0].price, 64)
	newBestAskPrice, _ := strconv.ParseFloat(m.orderBookManager.books[newKey].asks[0].price, 64)
	for _, market := range m.orderBookManager.markets {
		if market != newMarket {
			key := market + ":" + pair
			bestBidPrice, _ := strconv.ParseFloat(m.orderBookManager.books[key].bids[0].price, 64)
			bestAskPrice, _ := strconv.ParseFloat(m.orderBookManager.books[key].asks[0].price, 64)
			if bestBidPrice > newBestAskPrice {
				bidFee := bestBidPrice * 0.0004
				askFee := newBestAskPrice * 0.0004
				gapRate := bestBidPrice - newBestAskPrice - bidFee - askFee
				if gapRate > 0 {
					fmt.Println("bingo*********")
					// fmt.Println("the rate is %10.6f", gapRate)
				}
			}
			if bestAskPrice < newBestBidPrice {
				bidFee := newBestBidPrice * 0.0004
				askFee := bestAskPrice * 0.0004
				gapRate := newBestBidPrice - bestAskPrice - bidFee - askFee
				if gapRate > 0 {
					fmt.Println("bingo@@@@@@@@@@")
					// fmt.Println("the rate is %10.6f", gapRate)
				}
			}

		}
	}
	// currentMarket = event.market
	// for market in orderBookManager.markets
	//  if there are opportunity do arbitrage.
	// m.orderbookUpdateEvents <- orderbook
}

// func (m *Arbitrator) simulateTrade(first, second *OrderBook) bool {
// 	firstBidPrice, _ := strconv.ParseFloat(first.bids[0].price, 64)
// 	firstAskPrice, _ := strconv.ParseFloat(second.bids[0].price, 64)

// }

//Run start an go thread to find opportunity.
// func (m *Arbitrator) Run() error {
// 	for event := range m.orderbookUpdateEvents {
// 		fmt.Println("coming from arbitrator..")
// 		fmt.Println(event.market, event.symbol)
// 	}
// 	return nil
// }
