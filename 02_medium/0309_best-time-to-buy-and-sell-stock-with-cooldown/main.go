package main

// LeetCode #309: Best Time to Buy and Sell Stock with Cooldown
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	hold := -prices[0]
	cool := 0
	sell := 0

	for i := 1; i < len(prices); i++ {
		prevSell := sell
		sell = max(sell, hold+prices[i])
		hold = max(hold, cool-prices[i])
		cool = max(cool, prevSell)
	}

	return max(sell, cool)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
	fmt.Println(maxProfit([]int{1, 2, 4}))
}
