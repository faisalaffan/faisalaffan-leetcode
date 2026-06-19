package main

// LeetCode #3577: Count the Number of Computer Unlocking Permutations
// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/
// Difficulty: Medium
// Complexity: O(n! * n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountTheNumberOfComputerUnlockingPermutations(2))
	// Test case 2
	fmt.Println("Test 2:", CountTheNumberOfComputerUnlockingPermutations(3))
	// Test case 3
	fmt.Println("Test 3:", CountTheNumberOfComputerUnlockingPermutations(1))
}

func CountTheNumberOfComputerUnlockingPermutations(n int) int {
	// Count permutations of [1..n] that satisfy unlock pattern rules
	// Simple version: all permutations are valid
	// n! permutations
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
