package main

// LeetCode #638: Shopping Offers
// https://leetcode.com/problems/shopping-offers/
// Difficulty: Medium
// Time: O(n * special * states) where states is up to product of (needs[i]+1)
// Space: O(product of needs[i]+1) for memoization

import (
	"fmt"
)

func main() {
	fmt.Println(shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	fmt.Println(shoppingOffers([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	memo := make(map[string]int)
	return dfs(price, special, needs, n, memo)
}

func dfs(price []int, special [][]int, needs []int, n int, memo map[string]int) int {
	key := fmt.Sprint(needs)
	if val, ok := memo[key]; ok {
		return val
	}

	cost := 0
	for i := 0; i < n; i++ {
		cost += price[i] * needs[i]
	}

	for _, offer := range special {
		valid := true
		newNeeds := make([]int, n)
		for i := 0; i < n; i++ {
			if offer[i] > needs[i] {
				valid = false
				break
			}
			newNeeds[i] = needs[i] - offer[i]
		}
		if valid {
			cost = min(cost, offer[n]+dfs(price, special, newNeeds, n, memo))
		}
	}

	memo[key] = cost
	return cost
}
