package main

// LeetCode #2040: Kth Smallest Product of Two Sorted Arrays
// https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/
// Difficulty: Hard
//
// Approach: Binary search on the answer value. Count how many pairs have
// product <= mid using two-pointer technique (since both arrays are sorted).

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(kthSmallestProduct([]int{2, 5}, []int{3, 4}, 2))

	// Example 2
	fmt.Println(kthSmallestProduct([]int{-4, -2, 0, 3}, []int{2, 4}, 6))

	// Example 3
	fmt.Println(kthSmallestProduct([]int{-2, -1, 0, 1, 2}, []int{-3, -1, 2, 4, 5}, 3))
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int) int64 {
	left, right := int64(-1_000_000_000_000_000_000), int64(1_000_000_000_000_000_000)

	for left < right {
		mid := left + (right-left)/2
		if countLessOrEqual(nums1, nums2, mid) >= int64(k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// countLessOrEqual counts number of pairs (i,j) such that nums1[i]*nums2[j] <= x
func countLessOrEqual(nums1, nums2 []int, x int64) int64 {
	var count int64
	for _, a := range nums1 {
		if a > 0 {
			limit := x / int64(a)
			count += int64(sort.Search(len(nums2), func(i int) bool { return nums2[i] > int(limit) }))
		} else if a < 0 {
			limit := ceilDiv(x, int64(a))
			idx := sort.Search(len(nums2), func(i int) bool { return nums2[i] >= int(limit) })
			count += int64(len(nums2) - idx)
		} else if x >= 0 {
			count += int64(len(nums2))
		}
	}
	return count
}

// ceilDiv returns ceil(a/b) for possibly negative values. Assumes b != 0.
func ceilDiv(a, b int64) int64 {
	if a*b >= 0 {
		return (a + b - 1) / b
	}
	return a / b
}
