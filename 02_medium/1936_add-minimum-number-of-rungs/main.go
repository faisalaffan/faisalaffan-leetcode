package main

// LeetCode #1936: Add Minimum Number of Rungs
// https://leetcode.com/problems/add-minimum-number-of-rungs/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AddRungs([]int{1, 3, 5, 10}, 2))
	fmt.Println(AddRungs([]int{3, 6, 8, 10}, 3))
	fmt.Println(AddRungs([]int{3, 4, 6, 7}, 2))
}

// Time: O(n), Space: O(1)
func AddRungs(rungs []int, dist int) int {
	count := 0
	prev := 0
	for _, r := range rungs {
		gap := r - prev
		if gap > dist {
			count += (gap - 1) / dist
		}
		prev = r
	}
	return count
}
