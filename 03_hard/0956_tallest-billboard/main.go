package main

// LeetCode #956: Tallest Billboard
// https://leetcode.com/problems/tallest-billboard/
// Difficulty: Hard

import "fmt"

func tallestBillboard(rods []int) int {
	// dp[diff] = max total sum of both sides with this diff (left - right)
	dp := map[int]int{0: 0}

	for _, r := range rods {
		cur := make(map[int]int)
		for diff, total := range dp {
			// 1. skip this rod
			if total > cur[diff] {
				cur[diff] = total
			}
			// 2. add to left side (diff increases)
			left := total + r
			if left > cur[diff+r] {
				cur[diff+r] = left
			}
			// 3. add to right side (diff decreases)
			right := total + r
			newDiff := diff - r
			if newDiff < 0 {
				newDiff = -newDiff
			}
			if right > cur[newDiff] {
				cur[newDiff] = right
			}
		}
		dp = cur
	}

	// dp[0] is the max total sum with equal sides; each side is half
	return dp[0] / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	// Expected: 6

	fmt.Println("Example 2:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	// Expected: 10

	fmt.Println("Example 3:")
	fmt.Println(tallestBillboard([]int{1, 2}))
	// Expected: 0
}
