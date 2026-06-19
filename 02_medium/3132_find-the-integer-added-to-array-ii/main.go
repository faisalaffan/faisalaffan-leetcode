package main

// LeetCode #3132: Find the Integer Added to Array II
// https://leetcode.com/problems/find-the-integer-added-to-array-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	// Try all pairs from nums1 as the two removed elements
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			// Check if nums2 can be matched after removing nums1[i] and nums1[j]
			diff := -1001
			idx := 0
			match := true
			for k := 0; k < len(nums1) && match; k++ {
				if k == i || k == j {
					continue
				}
				curDiff := nums2[idx] - nums1[k]
				if diff == -1001 {
					diff = curDiff
				} else if curDiff != diff {
					match = false
				}
				idx++
			}
			if match && diff >= 0 {
				return diff
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10})) // Expected: -2
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))             // Expected: 2
}
