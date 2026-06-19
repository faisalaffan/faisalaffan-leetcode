package main

// LeetCode #2357: Make Array Zero by Subtracting Equal Amounts
// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// Difficulty: Easy
// Time O(n log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{1, 5, 0, 3, 5})) // 3
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{0}))              // 0
}

func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count++
			sub := nums[i]
			for j := i; j < len(nums); j++ {
				nums[j] -= sub
			}
		}
	}
	return count
}
