package main

// LeetCode #2972: Count the Number of Incremovable Subarrays II
// https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-ii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	j := n - 1
	for j > 0 && nums[j-1] < nums[j] {
		j--
	}
	if j == 0 {
		return int64(n * (n + 1) / 2)
	}
	ans := int64(n - j + 1)
	prev := math.MinInt
	for _, x := range nums {
		if x <= prev {
			break
		}
		prev = x
		for j < n && nums[j] <= x {
			j++
		}
		ans += int64(n - j + 1)
	}
	return ans
}

func main() {
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
	fmt.Println(incremovableSubarrayCount([]int{6, 5, 4, 3}))
}
