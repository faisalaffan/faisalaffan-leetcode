package main

// LeetCode #565: Array Nesting
// https://leetcode.com/problems/array-nesting/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (reuses input array as visited marker)

import "fmt"

func main() {
	fmt.Println(ArrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
	fmt.Println(ArrayNesting([]int{0, 1, 2}))
}

func ArrayNesting(nums []int) int {
	maxLen := 0
	visited := make([]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		count := 0
		cur := i
		for !visited[cur] {
			visited[cur] = true
			cur = nums[cur]
			count++
		}
		if count > maxLen {
			maxLen = count
		}
	}

	return maxLen
}
