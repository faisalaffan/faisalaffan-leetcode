package main

// LeetCode #2952: Minimum Number of Coins to be Added
// https://leetcode.com/problems/minimum-number-of-coins-to-be-added/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAddedCoins([]int{1, 4, 10}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 4, 10, 5, 7, 19}, 19))
	fmt.Println(minimumAddedCoins([]int{1, 1, 1}, 20))
}

func minimumAddedCoins(coins []int, target int) (ans int) {
	sort.Ints(coins)
	for i, s := 0, 1; s <= target; {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s <<= 1
			ans++
		}
	}
	return
}
