package main

// LeetCode #3855: Sum of K-Digit Numbers in a Range
// https://leetcode.com/problems/sum-of-k-digit-numbers-in-a-range/
// Difficulty: Hard
//
// Sum all numbers in range [l, r] that have exactly k digits.
// Numbers with k digits are in [10^(k-1), 10^k - 1] for k > 0.
// For k = 1, range is [0, 9].
//
// Approach: Compute intersection of [l, r] with k-digit range.
// Use arithmetic series sum formula.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfNumbers(1, 100, 2))
	// Example 2
	fmt.Println(sumOfNumbers(10, 25, 2))
	// Edge: k = 1
	fmt.Println(sumOfNumbers(0, 15, 1))
	// Edge: no overlap
	fmt.Println(sumOfNumbers(100, 200, 1))
}

func sumOfNumbers(l int, r int, k int) int {
	const mod = 1000000007

	start := 0
	end := 0
	if k == 1 {
		start = 0
		end = 9
	} else {
		start = 1
		for i := 1; i < k; i++ {
			start *= 10
		}
		end = start*10 - 1
	}

	if start < l {
		start = l
	}
	if end > r {
		end = r
	}
	if start > end {
		return 0
	}

	count := end - start + 1
	sum := int64(start+end) * int64(count) / 2
	return int(sum % int64(mod))
}
