package main

// LeetCode #3618: Split Array by Prime Indices
// https://leetcode.com/problems/split-array-by-prime-indices/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SplitArrayByPrimeIndices([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", SplitArrayByPrimeIndices([]int{10, 20, 30, 40}))
	// Test case 3
	fmt.Println("Test 3:", SplitArrayByPrimeIndices([]int{1, 2}))
}

func SplitArrayByPrimeIndices(nums []int) [][]int {
	var result [][]int
	var current []int
	for i, v := range nums {
		current = append(current, v)
		if isPrime(i) || i == len(nums)-1 {
			result = append(result, current)
			current = nil
		}
	}
	return result
}
