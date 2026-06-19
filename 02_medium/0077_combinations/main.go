package main

// LeetCode #77: Combinations
// https://leetcode.com/problems/combinations/
// Difficulty: Medium

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		// Prune: if remaining numbers are not enough, skip
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]

	// Test case 2
	fmt.Println(combine(1, 1)) // [[1]]

	// Test case 3
	fmt.Println(combine(4, 4)) // [[1 2 3 4]]
}

// Time: O(C(n,k) * k) | Space: O(k)
