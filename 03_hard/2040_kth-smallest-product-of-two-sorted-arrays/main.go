package main

// LeetCode #2040: Kth Smallest Product of Two Sorted Arrays
// https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/
// Difficulty: Hard
// Approach: Binary Search + Counting

import (
	"fmt"
	"sort"
)

func kthSmallestProduct(nums1 []int, nums2 []int, k int) int64 {
	// Split each array into negatives, zeros, positives
	// Use binary search on answer
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
			// For positive a, we need b <= x/a
			limit := x / int64(a)
			// Count elements in nums2 <= limit
			count += int64(sort.Search(len(nums2), func(i int) bool { return nums2[i] > int(limit) }))
		} else if a < 0 {
			// For negative a, we need b >= ceil(x/a)
			// x/a rounds toward zero, so we need careful handling
			// x / a where a < 0: the inequality a*b <= x means b >= x/a (since dividing by negative flips)
			// But integer division in Go truncates toward zero.
			// For negative a: x/a truncates toward zero.
			// Need: b >= ceil(x/a) when a < 0
			// ceil division for possibly negative numbers
			limit := ceilDiv(x, int64(a))
			// Count elements in nums2 >= limit
			idx := sort.Search(len(nums2), func(i int) bool { return nums2[i] >= int(limit) })
			count += int64(len(nums2) - idx)
		} else {
			// a == 0, product is 0
			if x >= 0 {
				count += int64(len(nums2))
			}
		}
	}
	return count
}

// ceilDiv returns ceil(a/b) for possibly negative values.
// Assumes b != 0.
func ceilDiv(a, b int64) int64 {
	if a*b >= 0 {
		// Same sign or zero: ceil division
		return (a + b - 1) / b
	}
	// Different signs: regular division in Go truncates toward zero, which is ceil for negative
	return a / b
}

func main() {
	fmt.Println("2040. Kth Smallest Product of Two Sorted Arrays")

	// Example 1
	nums1 := []int{2, 5}
	nums2 := []int{3, 4}
	k := 2
	fmt.Printf("nums1=%v nums2=%v k=%d → %d (expected 8)\n",
		nums1, nums2, k, kthSmallestProduct(nums1, nums2, k))

	// Example 2
	nums1 = []int{-4, -2, 0, 3}
	nums2 = []int{2, 4}
	k = 6
	fmt.Printf("nums1=%v nums2=%v k=%d → %d (expected 0)\n",
		nums1, nums2, k, kthSmallestProduct(nums1, nums2, k))

	// Example 3
	nums1 = []int{-2, -1, 0, 1, 2}
	nums2 = []int{-3, -1, 2, 4, 5}
	k = 3
	fmt.Printf("nums1=%v nums2=%v k=%d → %d (expected -6)\n",
		nums1, nums2, k, kthSmallestProduct(nums1, nums2, k))
}
