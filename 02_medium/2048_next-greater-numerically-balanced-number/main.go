package main

// LeetCode #2048: Next Greater Numerically Balanced Number
// https://leetcode.com/problems/next-greater-numerically-balanced-number/
// Difficulty: Medium
// Time: O(n) where n is the number until we find it | Space: O(1)

import "fmt"

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if isBalanced(i) {
			return i
		}
	}
}

func isBalanced(n int) bool {
	digits := make([]int, 10)
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}
	for d := 1; d <= 9; d++ {
		if digits[d] > 0 && digits[d] != d {
			return false
		}
	}
	return digits[0] == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", nextBeautifulNumber(1))
	// Expected: 22

	// Test case 2
	fmt.Println("Test 2:", nextBeautifulNumber(1000))
	// Expected: 1333

	// Test case 3
	fmt.Println("Test 3:", nextBeautifulNumber(3000))
	// Expected: 3133
}
