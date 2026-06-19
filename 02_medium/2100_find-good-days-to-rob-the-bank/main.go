package main

// LeetCode #2100: Find Good Days to Rob the Bank
// https://leetcode.com/problems/find-good-days-to-rob-the-bank/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	if n < 2*time+1 {
		return []int{}
	}

	// left[i] = number of consecutive non-increasing days ending at i
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	// right[i] = number of consecutive non-decreasing days starting at i
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	result := []int{}
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", goodDaysToRobBank([]int{1, 1, 1, 1, 1}, 0))
	// Expected: [0, 1, 2, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", goodDaysToRobBank([]int{1, 2, 3, 4, 5, 6}, 2))
	// Expected: []
}
