package common

//Arbitrator is a strategy which move the coin from exchange which has high price
//to the one which has low price.
type Arbitrator struct {
	accountManager   *AccountManager
	orderBookManager *OrderBookManager
}

// Run start an go thread to find opportunity.
func (m *Arbitrator) Run() error {

	return nil
}
