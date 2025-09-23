package buysellprofit

/*
	in: []int prices
	out: maxProfit

	set profit to 0; buyPrice to first index
	Loop through prices
	- check the currentProfit for if you sell on that day.
	- If profit for selling on that day is greater than the saved maxProfit, save as maxProfit
	- If the current day's price is less than the buyDay price then set new buyDay to current price

	https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
*/

func maxProfit(prices []int) int {
	profit := 0
	buyPrice := prices[0]
	for _, price := range prices {
		currentProfit := price - buyPrice
		if currentProfit > profit {
			profit = currentProfit
		}
		if buyPrice > price {
			buyPrice = price
		}

	}
	return profit
}
