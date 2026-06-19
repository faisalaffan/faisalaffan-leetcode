package main

// LeetCode #2110: Number of Smooth Descent Periods of a Stock
// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	var result int64 = 1 // single element
	length := 1

	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			length++
		} else {
			length = 1
		}
		result += int64(length)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getDescentPeriods([]int{3, 2, 1, 4}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", getDescentPeriods([]int{8, 6, 7, 7}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", getDescentPeriods([]int{1}))
	// Expected: 1
}
