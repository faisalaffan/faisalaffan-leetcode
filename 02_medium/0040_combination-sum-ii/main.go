package main

// LeetCode #40: Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]

	// Test case 2
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5)) // [[1 2 2] [5]]
}

// Time: O(2^n) | Space: O(n)
