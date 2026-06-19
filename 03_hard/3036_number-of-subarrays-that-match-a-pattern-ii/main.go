package main

// LeetCode #3036: Number of Subarrays That Match a Pattern II
// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/
// Difficulty: Hard

import (
	"cmp"
	"fmt"
)

func countMatchingSubarrays(nums, pattern []int) int {
	m := len(pattern)
	arr := make([]int, 0, m+1+len(nums)-1)
	arr = append(arr, pattern...)
	arr = append(arr, 2)
	for i := 1; i < len(nums); i++ {
		arr = append(arr, cmp.Compare(nums[i], nums[i-1]))
	}
	n := len(arr)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min2(r-i+1, z[i-l])
		}
		for i+z[i] < n && arr[z[i]] == arr[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	ans := 0
	for i := m + 1; i < n; i++ {
		if z[i] == m {
			ans++
		}
	}
	return ans
}
func min2(a, b int) int { if a < b { return a }; return b }

func main() {
	fmt.Println(countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))
	fmt.Println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
}
