package main

// LeetCode #319: Bulb Switcher
// https://leetcode.com/problems/bulb-switcher/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	// Only bulbs at perfect square positions are toggled an odd number of times
	return int(math.Sqrt(float64(n)))
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", bulbSwitch(3))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", bulbSwitch(0))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", bulbSwitch(25))
	// Expected: 5
}
