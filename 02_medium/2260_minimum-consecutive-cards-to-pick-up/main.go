package main

// LeetCode #2260: Minimum Consecutive Cards to Pick Up
// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumCardPickup(cards []int) int {
	last := make(map[int]int)
	minLen := len(cards) + 1

	for i, c := range cards {
		if prev, ok := last[c]; ok {
			dist := i - prev + 1
			if dist < minLen {
				minLen = dist
			}
		}
		last[c] = i
	}

	if minLen > len(cards) {
		return -1
	}
	return minLen
}

func main() {
	// Test case 1
	fmt.Println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumCardPickup([]int{1, 0, 5, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumCardPickup([]int{1, 2, 3, 4, 5, 1}))
	// Expected: 6
}
