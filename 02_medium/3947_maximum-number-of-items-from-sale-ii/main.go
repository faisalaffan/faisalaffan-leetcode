package main

// LeetCode #3947: Maximum Number of Items From Sale II
// https://leetcode.com/problems/maximum-number-of-items-from-sale-ii/
// Difficulty: Medium
// Time: O(N log N) | Space: O(N) where N = len(items)
// Approach: Each purchased copy of item i gives at most 1 free copy of a
// different item j where factor_i | factor_j (at most once per ordered pair).
// So first out_degree[i] copies of item i each give 2 items (1 purchased
// + 1 free). Subsequent copies give 1 each. Sort bonus copies by price,
// greedily buy cheapest bonus copies, then buy cheapest regular copies.

import (
	"fmt"
	"sort"
)

func MaximumNumberOfItemsFromSaleIi(items [][]int, budget int) int {
	m := len(items)

	// Compute out_degree per item (how many j where i|j, j != i)
	outDeg := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j && items[j][0]%items[i][0] == 0 {
				outDeg[i]++
			}
		}
	}

	// Collect all bonus opportunities: each gives 2 copies at price price_i
	type bonus struct {
		price int
	}
	var bonuses []bonus
	for i := 0; i < m; i++ {
		for k := 0; k < outDeg[i]; k++ {
			bonuses = append(bonuses, bonus{price: items[i][1]})
		}
	}

	// Sort bonuses by price
	sort.Slice(bonuses, func(i, j int) bool {
		return bonuses[i].price < bonuses[j].price
	})

	// Find min price for regular copies
	minPrice := int(1e9 + 1)
	for _, it := range items {
		if it[1] < minPrice {
			minPrice = it[1]
		}
	}

	remaining := budget
	total := 0

	// Buy cheapest bonuses first
	for _, b := range bonuses {
		if remaining < b.price {
			break
		}
		// Bonus gives 2 copies for price_i
		// Only buy if better than 2 regular copies at minPrice
		if b.price < 2*minPrice {
			remaining -= b.price
			total += 2
		}
	}

	// Buy regular copies with remaining budget
	total += remaining / minPrice

	return total
}

func main() {
	// Example 1
	fmt.Println(MaximumNumberOfItemsFromSaleIi([][]int{{1, 6}, {2, 4}, {3, 5}}, 19)) // Expected: 5

	// Example 2
	fmt.Println(MaximumNumberOfItemsFromSaleIi([][]int{{2, 8}, {1, 10}, {6, 6}, {4, 12}, {5, 20}, {5, 17}}, 35)) // Expected: 7
}
