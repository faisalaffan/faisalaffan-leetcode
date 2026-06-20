package main

// LeetCode #2499: Minimum Total Cost to Make Arrays Unequal
// https://leetcode.com/problems/minimum-total-cost-to-make-arrays-unequal/
// Difficulty: Hard
//
// Given two arrays nums1 and nums2, you can swap elements at the same
// index. The cost of swapping at index i is i. Find minimum total cost
// to make nums1[i] != nums2[i] for all i, or return -1 if impossible.
//
// Approach: Count positions where nums1[i] == nums2[i]. Accumulate cost.
// The dominant value needs extra swaps from other positions.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumTotalCost([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	// Example 2
	fmt.Println(minimumTotalCost([]int{2, 1, 2}, []int{1, 2, 2}))
	// Example 3
	fmt.Println(minimumTotalCost([]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}))
	// Edge: already unequal
	fmt.Println(minimumTotalCost([]int{1, 2}, []int{3, 4}))
}

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	same, n := 0, len(nums1)
	cnt := make([]int, n+1)
	for i, a := range nums1 {
		b := nums2[i]
		if a == b {
			same++
			cnt[a]++
		}
	}

	var ans int64
	var m, lead int
	for i, v := range cnt {
		if t := v*2 - same; t > 0 {
			m = t
			lead = i
			break
		}
	}

	for i, a := range nums1 {
		if same > 0 && a == nums2[i] {
			ans += int64(i)
			same--
		}
	}

	for i, a := range nums1 {
		b := nums2[i]
		if m > 0 && a != b && a != lead && b != lead {
			ans += int64(i)
			m--
		}
	}
	if m > 0 {
		return -1
	}
	return ans
}
