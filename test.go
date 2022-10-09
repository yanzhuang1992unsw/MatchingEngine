package main

import (
	"fmt"
	"meetingEngine/engine"
)

func Test() {
	var order1 = engine.Order{Price: 100, Amount: 20, IsSeller: true, ID: 1}
	var order2 = engine.Order{Price: 1000, Amount: 20, IsSeller: true, ID: 2}
	var order3 = engine.Order{Price: 900, Amount: 20, IsSeller: true, ID: 3}
	var order4 = engine.Order{Price: 900, Amount: 200, IsSeller: false, ID: 4}

	var book engine.Book
	trade1 := book.ProcessTrade(order1)
	fmt.Println(trade1) // []

	trade2 := book.ProcessTrade(order2)
	fmt.Println(trade2) // []

	trade3 := book.ProcessTrade(order3)
	fmt.Println(trade3) // []

	fmt.Println(book) // should be same as {[] [{100 20 true 1} {900 20 true 3} {1000 20 true 2}]}

	trade4 := book.ProcessTrade(order4)
	fmt.Println(trade4) // should be same as [{4 1 100 20} {4 3 900 20}]

	fmt.Println(book) // should be same as {[{900 160 false 4}] [{1000 20 true 2}]}
}
