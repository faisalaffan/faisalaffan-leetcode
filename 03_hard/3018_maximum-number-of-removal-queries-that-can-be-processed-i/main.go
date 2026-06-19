package main

// LeetCode #3018: Maximum Number of Removal Queries That Can Be Processed I
// https://leetcode.com/problems/maximum-number-of-removal-queries-that-can-be-processed-i/
// Difficulty: Hard [Paid]

import "fmt"

func maximumProcessableQueries(nums []int, queries []int) int {
	n := len(nums)
	m := len(queries)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		if m > 0 && nums[i] >= queries[0] {
			dp[i][i] = 1
		} else {
			dp[i][i] = 0
		}
	}
	ans := 0
	for length := 2; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			val := 0
			if l > 0 {
				prev := dp[l-1][r]
				if prev > val { val = prev }
				if prev >= 0 && prev < m && nums[l-1] >= queries[prev] {
					if prev+1 > val { val = prev + 1 }
				}
			}
			if r+1 < n {
				prev := dp[l][r+1]
				if prev > val { val = prev }
				if prev >= 0 && prev < m && nums[r+1] >= queries[prev] {
					if prev+1 > val { val = prev + 1 }
				}
			}
			dp[l][r] = val
			if val > ans { ans = val }
		}
	}
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			val := dp[l][r]
			if val > ans { ans = val }
			if l > 0 && val < m && nums[l-1] >= queries[val] && val+1 > ans { ans = val + 1 }
			if r+1 < n && val < m && nums[r+1] >= queries[val] && val+1 > ans { ans = val + 1 }
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumProcessableQueries([]int{1, 2, 3, 4}, []int{1, 2}))
}
