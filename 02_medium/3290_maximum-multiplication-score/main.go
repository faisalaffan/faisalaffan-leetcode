package main

// LeetCode #3290: Maximum Multiplication Score
// https://leetcode.com/problems/maximum-multiplication-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))          // 70
	fmt.Println(maxScore([]int{-1, -2, -3, -4}, []int{1, 2, 3, 4}))      // -20
	fmt.Println(maxScore([]int{3, 2, 1, 4}, []int{2, 3, 4, 5, 6}))       // 52
}

func maxScore(a []int, b []int) int64 {
	const negInf int64 = -1e18
	dp := [4]int64{negInf, negInf, negInf, negInf}

	for _, bi := range b {
		for i := 3; i >= 0; i-- {
			var prev int64
			if i > 0 {
				prev = dp[i-1]
			}
			val := prev + int64(a[i])*int64(bi)
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[3]
}
