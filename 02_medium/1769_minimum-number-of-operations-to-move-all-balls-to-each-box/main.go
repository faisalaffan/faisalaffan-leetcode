package main

// LeetCode #1769: Minimum Number of Operations to Move All Balls to Each Box
// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output

import "fmt"

func minOperations(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)

	// Left to right: count balls and accumulate moves
	balls := 0
	moves := 0
	for i := 0; i < n; i++ {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	// Right to left
	balls = 0
	moves = 0
	for i := n - 1; i >= 0; i-- {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	return result
}

func main() {
	fmt.Println(minOperations("110"))    // Expected: [1, 1, 3]
	fmt.Println(minOperations("001011")) // Expected: [11, 8, 5, 4, 3, 4]
	fmt.Println(minOperations("0"))      // Expected: [0]
}
