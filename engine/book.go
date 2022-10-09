package engine

type Book struct {
	// BuyerOrderList and SellerOrderList Ensure the list is sorted by price and time
	BuyerOrderList  []Order
	SellerOrderList []Order
}

func (book *Book) AddSellerList(sellerOrder Order) {
	// scan the list and find the right position
	for i, sellerOrderListItem := range book.SellerOrderList {

		// 1, if the new order's price is bigger than current one, add new order before the current one
		// 2, if new order's price is equal or smaller than current one, skip.
		// As if two orders' price are equal, the new order must be latter than the current one.
		// Therefore, skip
		if sellerOrder.Price < sellerOrderListItem.Price {
			book.SellerOrderList = append(book.SellerOrderList[:i+1], book.SellerOrderList[i:]...)
			book.SellerOrderList[i] = sellerOrder
			return
		}
	}

	// if the seller list is empty add to the list directly
	// or the new order's price is smallest in the list
	// add the order to the end of the list
	book.SellerOrderList = append(book.SellerOrderList, sellerOrder)
	return
}

func (book *Book) AddBuyerList(buyerOrder Order) {
	// scan the list and find the right position
	for i, buyerOrderListItem := range book.BuyerOrderList {

		// 1, if the new order's price is bigger than current one, add new order before the current one
		// 2, if new order's price is equal or smaller than current one, skip.
		// As if two orders' price are equal, the new order must be latter than the current one.
		// Therefore, skip
		if buyerOrder.Price > buyerOrderListItem.Price {
			book.BuyerOrderList = append(book.BuyerOrderList[:i+1], book.BuyerOrderList[i:]...)
			book.BuyerOrderList[i] = buyerOrder
			return
		}
	}

	// if the seller list is empty add to the list directly
	// or the new order's price is smallest in the list
	// add the order to the end of the list
	book.BuyerOrderList = append(book.BuyerOrderList, buyerOrder)
	return
}

func (book *Book) RemoveBuyerOrder(index int) {
	if index == len(book.BuyerOrderList)-1 {
		book.BuyerOrderList = book.BuyerOrderList[:index]
		return
	}
	book.BuyerOrderList = append(book.BuyerOrderList[:index], book.BuyerOrderList[index+1:]...)
	return
}

func (book *Book) RemoveSellerOrder(index int) {
	if index == len(book.SellerOrderList)-1 {
		book.SellerOrderList = book.SellerOrderList[:index]
		return
	}
	book.SellerOrderList = append(book.SellerOrderList[:index], book.SellerOrderList[index+1:]...)
	return
}
