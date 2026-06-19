package main

// LeetCode #2968: Apply Operations to Maximize Frequency Score
// https://leetcode.com/problems/apply-operations-to-maximize-frequency-score/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func maxFrequencyScore(nums []int, k int64) int {
	sort.Ints(nums)
	n := len(nums)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}
	ans := 1
	left := 0
	for right := 0; right < n; right++ {
		for left < right {
			mid := (left + right) / 2
			leftCost := int64(nums[mid])*int64(mid-left) - (prefix[mid] - prefix[left])
			rightCost := (prefix[right+1] - prefix[mid+1]) - int64(nums[mid])*int64(right-mid)
			if leftCost+rightCost <= k {
				break
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFrequencyScore([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(maxFrequencyScore([]int{1, 4, 4, 2, 7}, 6))
}
