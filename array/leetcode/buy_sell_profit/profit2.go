package buysellprofit

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150
	in: list of prices
	out: Max profit
	constraints: Able to buy and sell multiple times on the same day.

	Buy during valleys, sell during peaks.
	Compare the prices between days
	- if it's positive, add to sum
	- if it's negative, set the buy price to the new day
	O(n) time; O(1) space
*/

import (
	"fmt"
)

func maxProfit2(prices []int) int {
	sum := 0
	buyPrice := prices[0]
	for _, price := range prices {
		current := price - buyPrice
		if current > 0 {
			sum += current
		}
		buyPrice = price
		fmt.Printf("Sum: %v\t current: %v\n", sum, current)
	}
	return sum
}
