package main

// LeetCode #400: Nth Digit
// https://leetcode.com/problems/nth-digit/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	// 1-9: 9 digits,  10-99: 90*2 digits,  100-999: 900*3 digits
	length := 1
	count := 9
	start := 1

	for n > length*count {
		n -= length * count
		length++
		count *= 10
		start *= 10
	}

	// Find the actual number
	num := start + (n-1)/length
	// Find the digit within the number
	digitIdx := (n - 1) % length
	digitStr := strconv.Itoa(num)
	return int(digitStr[digitIdx] - '0')
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findNthDigit(3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", findNthDigit(11))
	// Expected: 0 (from 10)

	// Test case 3
	fmt.Println("Test 3:", findNthDigit(190))
	// Expected: 1 (from 100)
}
