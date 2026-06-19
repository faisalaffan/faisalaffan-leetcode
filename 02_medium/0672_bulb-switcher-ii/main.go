package main

// LeetCode #672: Bulb Switcher II
// https://leetcode.com/problems/bulb-switcher-ii/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(flipLights(1, 1))
	fmt.Println(flipLights(2, 1))
	fmt.Println(flipLights(3, 1))
}

func flipLights(n int, m int) int {
	// Only first 3 bulbs matter (key insight)
	if m == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if m == 1 {
			return 3
		}
		return 4
	}
	// n >= 3
	if m == 1 {
		return 4
	}
	if m == 2 {
		return 7
	}
	return 8
}
