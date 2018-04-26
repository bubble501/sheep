package common

//Arbitrator is a strategy which move the coin from exchange which has high price
//to the one which has low price.
type Arbitrator struct {
	accountManager   *AccountManager
	orderBookManager *OrderBookManager
}

func(m *Arbitrator) handleNewOrder(orderbook *JSONAddOrderBook) {
	symbol := orderbook.symbol
	market := orderbook.market
	m.orderBookManager.books;
}

// Run start an go thread to find opportunity.
func (m *Arbitrator) Run() error {

	return nil
}
