package main

// LeetCode #50: Pow(x, n)
// https://leetcode.com/problems/powx-n/
// Difficulty: Medium

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(myPow(2.0, 10)) // 1024

	// Test case 2
	fmt.Println(myPow(2.1, 3)) // 9.261

	// Test case 3
	fmt.Println(myPow(2.0, -2)) // 0.25
}

// Time: O(log n) | Space: O(1)
