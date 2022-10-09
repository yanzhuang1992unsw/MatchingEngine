package engine

type Trade struct {
	TradeOrderId uint64
	OrderId      uint64
	Price        float32
	Amount       uint64
}

func (book *Book) ProcessTrade(order Order) []Trade {
	if order.IsSeller {
		return book.processSellerTrade(order)
	}
	return book.processBuyerTrade(order)
}

func (book *Book) processBuyerTrade(order Order) (trades []Trade) {
	// if the list is not empty and the smallest price of seller is bigger than the buying order
	if len(book.SellerOrderList) != 0 && book.SellerOrderList[0].Price <= order.Price {

		// deal with the finished order
		var removedItemList []int
		for index, sellerOrderListItem := range book.SellerOrderList {
			if sellerOrderListItem.Price > order.Price {
				for i := 0; i < len(removedItemList); i++ {
					book.RemoveSellerOrder(removedItemList[i] - i)
				}
				break
			}

			// if the seller order's amount is more than order's amount
			if sellerOrderListItem.Amount >= order.Amount {
				trades = append(trades, Trade{
					order.ID,
					sellerOrderListItem.ID,
					sellerOrderListItem.Price,
					order.Amount})
				sellerOrderListItem.Amount -= order.Amount
				if sellerOrderListItem.Amount == 0 {
					removedItemList = append(removedItemList, index)
				}
				return
			}

			// if the seller order's amount is shorter than order's amount
			if sellerOrderListItem.Amount < order.Amount {
				trades = append(trades, Trade{
					order.ID,
					sellerOrderListItem.ID,
					sellerOrderListItem.Price,
					sellerOrderListItem.Amount})
				order.Amount -= sellerOrderListItem.Amount
				removedItemList = append(removedItemList, index)
				continue
			}
		}
	}
	book.AddBuyerList(order)
	return
}

func (book *Book) processSellerTrade(order Order) (trades []Trade) {
	// if the list is not empty and the biggest price of buyer is bigger than the selling order
	if len(book.BuyerOrderList) != 0 && book.BuyerOrderList[0].Price >= order.Price {

		// deal with the finished order
		var removedItemList []int
		for index, buyerOrderListItem := range book.BuyerOrderList {
			if buyerOrderListItem.Price < order.Price {
				for i := 0; i < len(removedItemList); i++ {
					book.RemoveBuyerOrder(removedItemList[i] - i)
				}
				break
			}

			// if the buyer order's amount is more than order's amount
			if buyerOrderListItem.Amount >= order.Amount {
				trades = append(trades, Trade{
					order.ID,
					buyerOrderListItem.ID,
					buyerOrderListItem.Price,
					order.Amount})
				buyerOrderListItem.Amount -= order.Amount
				if buyerOrderListItem.Amount == 0 {
					removedItemList = append(removedItemList, index)
				}
				return
			}

			// if the buyer's amount is shorter than order's amount
			if buyerOrderListItem.Amount < order.Amount {
				trades = append(trades, Trade{
					order.ID,
					buyerOrderListItem.ID,
					buyerOrderListItem.Price,
					buyerOrderListItem.Amount})
				order.Amount -= buyerOrderListItem.Amount
				removedItemList = append(removedItemList, index)
				continue
			}
		}
	}

	// if there is still remain or if there is no match add to the list
	book.AddSellerList(order)
	return
}
