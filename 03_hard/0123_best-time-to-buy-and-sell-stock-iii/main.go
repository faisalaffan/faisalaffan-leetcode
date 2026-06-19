package main

// LeetCode #123: Best Time to Buy and Sell Stock III
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("123. Best Time to Buy and Sell Stock III")
	fmt.Println("[3,3,5,0,0,3,1,4] ->", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}), "(expected 6)")
	fmt.Println("[1,2,3,4,5] ->", maxProfit([]int{1, 2, 3, 4, 5}), "(expected 4)")
	fmt.Println("[7,6,4,3,1] ->", maxProfit([]int{7, 6, 4, 3, 1}), "(expected 0)")
	fmt.Println("[1] ->", maxProfit([]int{1}), "(expected 0)")
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// State machine with at most 2 transactions:
	// buy1 = min price seen for first buy
	// sell1 = max profit after first sell
	// buy2 = effective cost after first profit (min price2 - profit1)
	// sell2 = max profit after second sell
	buy1 := math.MaxInt32
	buy2 := math.MaxInt32
	sell1 := 0
	sell2 := 0

	for _, price := range prices {
		if price < buy1 {
			buy1 = price
		}
		if price-buy1 > sell1 {
			sell1 = price - buy1
		}
		if price-sell1 < buy2 {
			buy2 = price - sell1
		}
		if price-buy2 > sell2 {
			sell2 = price - buy2
		}
	}

	return sell2
}
