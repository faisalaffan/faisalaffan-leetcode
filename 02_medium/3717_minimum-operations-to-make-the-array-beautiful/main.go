package main

// LeetCode #3717: Minimum Operations to Make the Array Beautiful
// https://leetcode.com/problems/minimum-operations-to-make-the-array-beautiful/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToMakeTheArrayBeautiful(nums []int) int {
	ops := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]%prev != 0 {
			target := ((nums[i] / prev) + 1) * prev
			ops += target - nums[i]
			prev = target
		} else {
			prev = nums[i]
		}
	}
	return ops
}

func main() {
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{3, 7, 9}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{1, 1, 1}))
	fmt.Println(minimumOperationsToMakeTheArrayBeautiful([]int{2, 3, 5}))
}
