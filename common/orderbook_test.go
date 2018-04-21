package common

// func createOrders(i int,  j int, bid bool) ([]Order) {
// 	orders := make([]Order, 10)
// 	for k:= i; k < j; k++ {
// 		index := i
// 		if bid==false {
// 			index = j-i
// 		}
// 		orders[k-i] = Order {
// 			price: float32(index)*100.0,
// 			volume:float32(index)*10.0,
// 		}
// 	}
// 	return orders
// }

// func TestOrderBookGetInstance(t *testing.T) {
// 	obm := GetOrderbookMap()
// 	assert.NotNil(t, obm)
// }

// func TestReplaceOrderbook(t *testing.T) {
// 	obm := GetOrderbookMap()
// 	ob := Orderbook {
// 		bids: make([]Order, 10),
// 		asks: make([]Order, 10),
// 	}
// 	ReplaceOrderbook("huobi", "ethusdt", ob)
// 	fetchedOb := (*obm)["huobi"]["ethusdt"]
// 	assert.Equal(t, ob, fetchedOb)
// }

// func TestMergeOrderbook(t *testing.T) {
// 	ob := Orderbook {
// 		bids: []Order{
// 			Order{
// 				price: 10.0,
// 				volume: 20.0,
// 			},
// 			Order{
// 				price: 30.0,
// 				volume: 40.0,
// 			},
// 		},
// 		asks: []Order{
// 			Order{
// 				price: 40.0,
// 				volume: 1.0,
// 			},
// 			Order{
// 				price: 30.0,
// 				volume: 40.0,
// 			},
// 		},
// 	}
// 	MergeOrderbook("huobi", "ethusdt", ob)
// 	fetchedOb := (*GetOrderbookMap())["huobi"]["ethusdt"]
// }
