package main

// LeetCode #1150: Check If a Number Is Majority Element in a Sorted Array
// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isMajorityElement([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5)) // true
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 101))        // false
}

// LeetCode submission: isMajorityElement
func isMajorityElement(nums []int, target int) bool {
	n := len(nums)
	// Binary search for first occurrence
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	first := lo
	last := first + n/2
	return last < n && nums[last] == target
}
