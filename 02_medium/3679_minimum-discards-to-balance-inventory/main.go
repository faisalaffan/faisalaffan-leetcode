package main

// LeetCode #3679: Minimum Discards to Balance Inventory
// https://leetcode.com/problems/minimum-discards-to-balance-inventory/
// Difficulty: Medium
// Time: O(n) | Space: O(max(arrivals))

import "fmt"

func minimumDiscardsToBalanceInventory(arrivals []int, w int, m int) int {
	maxVal := 0
	for _, v := range arrivals {
		if v > maxVal {
			maxVal = v
		}
	}

	cnt := make([]int, maxVal+1)
	ans := 0

	for i, x := range arrivals {
		if cnt[x] == m {
			arrivals[i] = 0
			ans++
		} else {
			cnt[x]++
		}

		left := i + 1 - w
		if left >= 0 {
			cnt[arrivals[left]]--
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3, 3, 3, 4}, 3, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 1, 1, 2, 2}, 2, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3}, 3, 1))
}
