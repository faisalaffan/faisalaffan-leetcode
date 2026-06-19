package main

// LeetCode #3693: Climbing Stairs II
// https://leetcode.com/problems/climbing-stairs-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func climbingStairsIi(n int, costs []int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for j := 0; j < n; j++ {
		cur := dp2 + costs[j] + 1
		if j >= 1 {
			cand := dp1 + costs[j] + 4
			if cand < cur {
				cur = cand
			}
		}
		if j >= 2 {
			cand := dp0 + costs[j] + 9
			if cand < cur {
				cur = cand
			}
		}
		dp0, dp1, dp2 = dp1, dp2, cur
	}
	return dp2
}

func main() {
	fmt.Println(climbingStairsIi(3, []int{1, 2, 3}))
	fmt.Println(climbingStairsIi(2, []int{5, 10}))
	fmt.Println(climbingStairsIi(5, []int{1, 1, 1, 1, 1}))
}
