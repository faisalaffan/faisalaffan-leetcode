package main

// LeetCode #2915: Length of the Longest Subsequence That Sums to Target
// https://leetcode.com/problems/length-of-the-longest-subsequence-that-sums-to-target/
// Difficulty: Medium
// Time: O(n*target) | Space: O(n*target)

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubsequence([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(lengthOfLongestSubsequence([]int{4, 1, 3, 2, 1, 5}, 7))
	fmt.Println(lengthOfLongestSubsequence([]int{1, 1, 5, 4, 5}, 3))
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	const negInf = -(1 << 30)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
		for j := range f[i] {
			f[i][j] = negInf
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		for j := 0; j <= target; j++ {
			f[i][j] = f[i-1][j]
			if j >= x && f[i-1][j-x]+1 > f[i][j] {
				f[i][j] = f[i-1][j-x] + 1
			}
		}
	}
	if f[n][target] <= 0 {
		return -1
	}
	return f[n][target]
}
