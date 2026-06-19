package main

// LeetCode #121: Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MaxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))
}
