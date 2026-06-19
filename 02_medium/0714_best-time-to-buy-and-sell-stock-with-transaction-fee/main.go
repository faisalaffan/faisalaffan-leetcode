package main

// LeetCode #714: Best Time to Buy and Sell Stock with Transaction Fee
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	cash := 0
	hold := -prices[0]

	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}

	return cash
}
