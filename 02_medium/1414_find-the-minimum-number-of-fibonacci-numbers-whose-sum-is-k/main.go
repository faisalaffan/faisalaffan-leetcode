package main

// LeetCode #1414: Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
// https://leetcode.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findMinFibonacciNumbers(7)) // 2

	// Test case 2
	fmt.Println(findMinFibonacciNumbers(10)) // 2

	// Test case 3
	fmt.Println(findMinFibonacciNumbers(19)) // 3

	// Test case 4
	fmt.Println(findMinFibonacciNumbers(1)) // 1
}

// Time: O(log k) since Fibonacci numbers grow exponentially
// Space: O(1)
func findMinFibonacciNumbers(k int) int {
	// Generate all Fibonacci numbers <= k
	fib := []int{1, 1}
	for fib[len(fib)-1] <= k {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}

	count := 0
	remaining := k
	for i := len(fib) - 1; i >= 0; i-- {
		if fib[i] <= remaining {
			remaining -= fib[i]
			count++
		}
		if remaining == 0 {
			break
		}
	}

	return count
}
