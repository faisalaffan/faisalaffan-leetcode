package main

// LeetCode #3562: Maximum Profit from Trading Stocks with Discounts
// https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts/
// Difficulty: Hard
//
// Given stock prices and a discount coupon that can be used once to buy
// a stock at half price, maximize profit from at most one buy-sell pair.
//
// Approach: Track minimum price with and without discount.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(maxProfitWithDiscount([]int{1, 2, 3, 4}, []int{2, 3, 4, 5}))
	// Example 2
	fmt.Println(maxProfitWithDiscount([]int{7, 1, 5, 3, 6, 4}, []int{7, 2, 5, 4, 7, 5}))
	// Edge: single day
	fmt.Println(maxProfitWithDiscount([]int{5}, []int{5}))
	// Edge: no profit possible
	fmt.Println(maxProfitWithDiscount([]int{5, 4, 3}, []int{5, 4, 3}))
}

func maxProfitWithDiscount(prices []int, discounted []int) int64 {
	n := len(prices)
	if n < 2 {
		return 0
	}

	minPrice := math.MaxInt32
	minDiscounted := math.MaxInt32
	var maxProfit int64

	for i := 0; i < n; i++ {
		// Sell at regular price
		if minPrice != math.MaxInt32 {
			profit := int64(prices[i] - minPrice)
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		// Sell having bought with discount
		if minDiscounted != math.MaxInt32 {
			profit := int64(prices[i] - minDiscounted)
			if profit > maxProfit {
				maxProfit = profit
			}
		}

		// Update min prices
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if discounted[i] < minDiscounted {
			minDiscounted = discounted[i]
		}
	}

	return maxProfit
}
