package main

// LeetCode #719: Find K-th Smallest Pair Distance
// https://leetcode.com/problems/find-k-th-smallest-pair-distance/
// Difficulty: Hard
//
// Algorithm: Binary Search + Two-Pointer Counting
// 1. Sort the array
// 2. Binary search on distance (0 to max-min)
// 3. Count pairs with distance <= mid using two-pointer
// 4. Find smallest distance with count >= k

import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	low, high := 0, nums[len(nums)-1]-nums[0]

	for low < high {
		mid := low + (high-low)/2
		count := countPairs(nums, mid)
		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// countPairs counts number of pairs with distance <= maxDist using two-pointer
func countPairs(nums []int, maxDist int) int {
	count := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > maxDist {
			left++
		}
		count += right - left
	}
	return count
}

func main() {
	// Example from problem
	nums1 := []int{1, 3, 1}
	k1 := 1
	result1 := smallestDistancePair(nums1, k1)
	fmt.Printf("Input: nums=%v, k=%d\nOutput: %d (expected: 0)\n\n", nums1, k1, result1)

	// Test case 2
	nums2 := []int{1, 1, 1}
	k2 := 2
	result2 := smallestDistancePair(nums2, k2)
	fmt.Printf("Input: nums=%v, k=%d\nOutput: %d (expected: 0)\n\n", nums2, k2, result2)

	// Test case 3
	nums3 := []int{1, 6, 1, 3, 4}
	k3 := 4
	result3 := smallestDistancePair(nums3, k3)
	fmt.Printf("Input: nums=%v, k=%d\nOutput: %d\n\n", nums3, k3, result3)

	// Test case 4: larger example
	nums4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k := 1; k <= 5; k++ {
		result := smallestDistancePair(nums4, k)
		fmt.Printf("k=%d -> %d\n", k, result)
	}
	fmt.Println()

	// Test case 5: negative numbers
	nums5 := []int{-5, -1, 0, 2, 7}
	k5 := 3
	result5 := smallestDistancePair(nums5, k5)
	fmt.Printf("Input: nums=%v, k=%d\nOutput: %d\n", nums5, k5, result5)
}
