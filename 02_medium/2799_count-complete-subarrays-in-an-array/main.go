package main

// LeetCode #2799: Count Complete Subarrays in an Array
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountCompleteSubarraysInAnArray(nums []int) int {
	// Count distinct elements
	distinct := make(map[int]bool)
	for _, n := range nums {
		distinct[n] = true
	}
	target := len(distinct)

	left := 0
	count := 0
	freq := make(map[int]int)
	unique := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			unique++
		}

		for unique == target {
			// All subarrays from left to right-end are valid
			count += len(nums) - right
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				unique--
			}
			left++
		}
	}

	return count
}

func main() {
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 3, 1, 2, 2}))
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 1}))
}
