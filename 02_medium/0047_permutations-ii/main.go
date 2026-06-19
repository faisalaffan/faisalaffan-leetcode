package main

// LeetCode #47: Permutations II
// https://leetcode.com/problems/permutations-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == n {
			perm := make([]int, n)
			copy(perm, path)
			result = append(result, perm)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, n))
	return result
}

func main() {
	// Test case 1
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1 1 2] [1 2 1] [2 1 1]]

	// Test case 2
	fmt.Println(permuteUnique([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}

// Time: O(n * n!) | Space: O(n)
