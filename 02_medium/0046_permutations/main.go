package main

// LeetCode #46: Permutations
// https://leetcode.com/problems/permutations/
// Difficulty: Medium

import "fmt"

func permute(nums []int) [][]int {
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
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]

	// Test case 2
	fmt.Println(permute([]int{0, 1})) // [[0 1] [1 0]]

	// Test case 3
	fmt.Println(permute([]int{1})) // [[1]]
}

// Time: O(n * n!) | Space: O(n)
