package main

// LeetCode #2310: Sum of Numbers With Units Digit K
// https://leetcode.com/problems/sum-of-numbers-with-units-digit-k/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if (i*k)%10 == num%10 && i*k <= num {
			return i
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println(minimumNumbers(58, 9))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumNumbers(37, 2))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumNumbers(0, 7))
	// Expected: 0
}
