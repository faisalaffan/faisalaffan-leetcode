package main

// LeetCode #78: Subsets
// https://leetcode.com/problems/subsets/
// Difficulty: Medium

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		n := len(result)
		for i := 0; i < n; i++ {
			newSubset := make([]int, len(result[i])+1)
			copy(newSubset, result[i])
			newSubset[len(result[i])] = num
			result = append(result, newSubset)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(subsets([]int{1, 2, 3}))
	// [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]

	// Test case 2
	fmt.Println(subsets([]int{0})) // [[] [0]]

	// Test case 3
	fmt.Println(subsets([]int{})) // [[]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
