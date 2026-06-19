package main

// LeetCode #2517: Maximum Tastiness of Candy Basket
// https://leetcode.com/problems/maximum-tastiness-of-candy-basket/
// Difficulty: Medium
// Time: O(n log n + n log max) | Space: O(1)
// Binary search on answer. Check: can we pick k candies with min diff >= mid?

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)) // 8
	fmt.Println(maximumTastiness([]int{1, 3, 1}, 2))             // 2
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	lo, hi := 0, price[len(price)-1]-price[0]

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canPick(price, k, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func canPick(price []int, k, minDiff int) bool {
	count, last := 1, price[0]
	for i := 1; i < len(price); i++ {
		if price[i]-last >= minDiff {
			count++
			last = price[i]
		}
	}
	return count >= k
}
