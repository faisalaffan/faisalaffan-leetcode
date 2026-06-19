package main

// LeetCode #2165: Smallest Value of the Rearranged Number
// https://leetcode.com/problems/smallest-value-of-the-rearranged-number/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}

	negative := num < 0
	if negative {
		num = -num
	}

	// Extract digits
	digits := []int{}
	for num > 0 {
		digits = append(digits, int(num%10))
		num /= 10
	}

	sort.Ints(digits)

	if negative {
		// Largest number from digits (descending)
		result := int64(0)
		for i := len(digits) - 1; i >= 0; i-- {
			result = result*10 + int64(digits[i])
		}
		return -result
	}

	// Smallest number: place smallest non-zero digit first
	// Find first non-zero digit
	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}
	// Swap first non-zero with position 0
	if i < len(digits) {
		digits[0], digits[i] = digits[i], digits[0]
	}

	result := int64(0)
	for _, d := range digits {
		result = result*10 + int64(d)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", smallestNumber(310))
	// Expected: 103

	// Test case 2
	fmt.Println("Test 2:", smallestNumber(-7605))
	// Expected: -7650

	// Test case 3
	fmt.Println("Test 3:", smallestNumber(0))
	// Expected: 0
}
