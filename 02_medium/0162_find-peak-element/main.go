package main

// LeetCode #162: Find Peak Element
// https://leetcode.com/problems/find-peak-element/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1}))
}
