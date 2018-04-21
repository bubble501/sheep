package common

//OrderBookEvent will be generate when an orderbook has been updated.
type OrderBookEvent struct {
	market string
	symbol string
}

//OrderBookEventHandler is a handler to handle OrderBookEvent.
type OrderBookEventHandler interface {
	handleEvent(event OrderBookEvent)
}
