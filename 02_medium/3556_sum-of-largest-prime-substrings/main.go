package main

// LeetCode #3556: Sum of Largest Prime Substrings
// https://leetcode.com/problems/sum-of-largest-prime-substrings/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import (
	"fmt"
	"strconv"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SumOfLargestPrimeSubstrings("237"))
	// Test case 2
	fmt.Println("Test 2:", SumOfLargestPrimeSubstrings("1234"))
	// Test case 3
	fmt.Println("Test 3:", SumOfLargestPrimeSubstrings("111"))
}

func SumOfLargestPrimeSubstrings(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s) && j-i <= 6; j++ { // at most 6 digits to avoid overflow
			num, _ := strconv.Atoi(s[i:j])
			if isPrime(num) {
				sum += num
			}
		}
	}
	return sum
}
