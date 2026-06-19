package main

// LeetCode #3134: Find the Median of the Uniqueness Array
// https://leetcode.com/problems/find-the-median-of-the-uniqueness-array/
// Difficulty: Hard
//
// The uniqueness array of nums is an array of distinct-counts for all subarrays.
// Find the median of this sorted uniqueness array.
// Binary search on the answer + sliding window count of subarrays with distinct count <= mid.

import (
	"fmt"
)

func main() {
	// Example: [1,2,3] -> 1
	nums := []int{1, 2, 3}
	fmt.Println(medianOfUniquenessArray(nums))

	// Additional test
	nums2 := []int{3, 4, 3, 4, 5}
	fmt.Println(medianOfUniquenessArray(nums2))
}

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	// The median is the (total+1)/2-th element (1-indexed) in sorted order
	medianPos := (total + 1) / 2

	// Binary search for the smallest x such that
	// count of subarrays with distinct count <= x >= medianPos
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if int(countLE(nums, mid)) >= medianPos {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// countLE returns the number of subarrays with distinct element count <= k
func countLE(nums []int, k int) int64 {
	n := len(nums)
	freq := make(map[int]int)
	distinct := 0
	var count int64 = 0
	left := 0

	for right := 0; right < n; right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			distinct++
		}
		for distinct > k {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinct--
			}
			left++
		}
		// All subarrays ending at 'right' with start in [left, right] have <= k distinct
		count += int64(right - left + 1)
	}
	return count
}
