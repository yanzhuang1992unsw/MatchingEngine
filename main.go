package main

import (
	"fmt"
	"meetingEngine/engine"
)

func main() {
	// Test()

	// init the variables
	var (
		book engine.Book
		id   uint64
	)

	// loop to waiting to the input
	for true {

		// order id
		id++
		var (
			price     float32
			amount    uint64
			isSelling bool
		)

		// input
		fmt.Printf("please input order `price`, `amount` and if it is selling (`true` or `false`) with space\n")
		_, err := fmt.Scanf("%f %d %t\n", &price, &amount, &isSelling)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		// format order
		var order = engine.Order{
			ID:       id,
			Price:    price,
			Amount:   amount,
			IsSeller: isSelling,
		}

		// get trade result
		trade := book.ProcessTrade(order)

		// if the trade is made than the result should be a list of trade with
		// {tradeID, orderID, price, amount}
		if trade != nil {
			fmt.Println(trade)
		}
	}
	//done := make(chan bool)
}
