package main

// LeetCode #152: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		result = max(result, maxProd)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
