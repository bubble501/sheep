package common

import "fmt"

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
	// currentMarket = event.market
	// for market in orderBookManager.markets
	//  if there are opportunity do arbitrage.
	// m.orderbookUpdateEvents <- orderbook
}

//Run start an go thread to find opportunity.
// func (m *Arbitrator) Run() error {
// 	for event := range m.orderbookUpdateEvents {
// 		fmt.Println("coming from arbitrator..")
// 		fmt.Println(event.market, event.symbol)
// 	}
// 	return nil
// }
