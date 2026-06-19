package main

// LeetCode #3497: Analyze Subscription Conversion
// https://leetcode.com/problems/analyze-subscription-conversion/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	conversions := []int{1, 2, 3, 4, 5}
	threshold := 3
	fmt.Println("Test 1:", AnalyzeSubscriptionConversion(conversions, threshold))

	// Test case 2
	conversions2 := []int{10, 20, 30, 5, 15}
	threshold2 := 10
	fmt.Println("Test 2:", AnalyzeSubscriptionConversion(conversions2, threshold2))

	// Test case 3
	conversions3 := []int{1}
	threshold3 := 1
	fmt.Println("Test 3:", AnalyzeSubscriptionConversion(conversions3, threshold3))
}

func AnalyzeSubscriptionConversion(conversions []int, threshold int) int {
	// Count users whose conversions exceed threshold
	sort.Ints(conversions)
	count := 0
	for _, c := range conversions {
		if c >= threshold {
			count++
		}
	}
	return count
}
