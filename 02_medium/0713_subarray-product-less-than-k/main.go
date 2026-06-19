package main

// LeetCode #713: Subarray Product Less Than K
// https://leetcode.com/problems/subarray-product-less-than-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	product := 1
	left := 0

	for right := 0; right < len(nums); right++ {
		product *= nums[right]

		for product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}
