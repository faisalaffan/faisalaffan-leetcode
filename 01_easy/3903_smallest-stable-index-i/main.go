package main

// LeetCode #3903: Smallest Stable Index I
// https://leetcode.com/problems/smallest-stable-index-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestStableIndexI([]int{5, 0, 1, 4}, 3))
	fmt.Println(SmallestStableIndexI([]int{3, 2, 1}, 1))
	fmt.Println(SmallestStableIndexI([]int{0}, 0))
}

// Time: O(n)
// Space: O(n)
func SmallestStableIndexI(nums []int, k int) int {
	n := len(nums)

	prefixMax := make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if prefixMax[i]-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}
