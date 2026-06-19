package main

// LeetCode #1838: Frequency of the Most Frequent Element
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	total := 0
	maxFreq := 0

	for right := 0; right < len(nums); right++ {
		total += nums[right]

		// Shrink window if we can't make all elements in window equal
		for nums[right]*(right-left+1)-total > k {
			total -= nums[left]
			left++
		}

		if right-left+1 > maxFreq {
			maxFreq = right - left + 1
		}
	}
	return maxFreq
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))       // Expected: 3
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))   // Expected: 2
	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))       // Expected: 1
}
