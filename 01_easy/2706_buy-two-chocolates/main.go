package main

// LeetCode #2706: Buy Two Chocolates
// https://leetcode.com/problems/buy-two-chocolates/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BuyTwoChocolates([]int{1, 2, 2}, 3))
	fmt.Println(BuyTwoChocolates([]int{3, 2, 3}, 3))
}

func BuyTwoChocolates(prices []int, money int) int {
	sort.Ints(prices)
	if len(prices) >= 2 && prices[0]+prices[1] <= money {
		return money - prices[0] - prices[1]
	}
	return money
}
