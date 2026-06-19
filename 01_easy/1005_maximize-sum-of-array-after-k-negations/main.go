package main

// LeetCode #1005: Maximize Sum Of Array After K Negations
// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))      // 5
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))  // 6
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)) // 13
}

// largestSumAfterKNegations maximizes the sum by negating k elements.
// Time: O(n log n). Space: O(1).
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	// Flip negatives
	for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i++ {
		nums[i] = -nums[i]
		k--
	}
	// If k is odd, flip the smallest absolute value
	if k%2 == 1 {
		minIdx := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] = -nums[minIdx]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
