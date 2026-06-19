package main

// LeetCode #1802: Maximum Value at a Given Index in a Bounded Array
// https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/
// Difficulty: Medium
// Time: O(log maxVal), Space: O(1)

import "fmt"

func maxValue(n int, index int, maxSum int) int {
	left, right := 1, maxSum

	for left < right {
		mid := left + (right-left+1)/2
		if canMake(n, index, maxSum, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canMake(n, index, maxSum, val int) bool {
	// Sum of left side (decreasing from val to some minimum)
	leftLen := index
	rightLen := n - index - 1

	total := val // the peak
	total += sumTriangle(val-1, leftLen)
	total += sumTriangle(val-1, rightLen)

	return total <= maxSum
}

// sum of sequence starting from max down to some minimum, limited by count
func sumTriangle(max, count int) int {
	if count <= 0 {
		return 0
	}
	if max >= count {
		// Enough height: use arithmetic series
		// max, max-1, ..., max-count+1
		return (max + max - count + 1) * count / 2
	}
	// Not enough height: max + (max-1) + ... + 1 + 1 + ... + 1 (count - max times)
	return max*(max+1)/2 + (count - max)
}

func main() {
	fmt.Println(maxValue(4, 2, 6))  // Expected: 2
	fmt.Println(maxValue(6, 1, 10)) // Expected: 3
	fmt.Println(maxValue(8, 7, 14)) // Expected: 4
}
