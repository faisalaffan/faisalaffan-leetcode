package main

// LeetCode #1703: Minimum Adjacent Swaps for K Consecutive Ones
// https://leetcode.com/problems/minimum-adjacent-swaps-for-k-consecutive-ones/
// Difficulty: Hard
// Strategy: Collect positions of 1s. For k consecutive 1s,
// use median prefix sum to compute min adjacent swaps.

import (
	"fmt"
	"math"
)

func minMoves(nums []int, k int) int {
	// Collect indices of 1s
	pos := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			pos = append(pos, i)
		}
	}

	n := len(pos)
	if n < k {
		return 0
	}

	// prefix sum of positions
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + pos[i]
	}

	ans := math.MaxInt32
	for i := 0; i <= n-k; i++ {
		// Window of k positions: pos[i..i+k-1]
		mid := i + k/2
		medianPos := pos[mid]

		// Number of elements on left side of median (in window)
		leftCount := mid - i
		// Number of elements on right side of median (in window)
		rightCount := i + k - 1 - mid

		// Sum of positions on left
		leftSum := pref[mid] - pref[i]
		// Sum of positions on right
		rightSum := pref[i+k] - pref[mid+1]

		// Ideal left positions: medianPos-1, medianPos-2, ...
		leftIdeal := leftCount*medianPos - leftCount*(leftCount+1)/2
		// Actual left sum
		leftActual := leftSum

		// Ideal right positions: medianPos+1, medianPos+2, ...
		rightIdeal := rightCount*medianPos + rightCount*(rightCount+1)/2
		// Actual right sum
		rightActual := rightSum

		swaps := (leftIdeal - leftActual) + (rightActual - rightIdeal)
		if swaps < ans {
			ans = swaps
		}
	}
	return ans
}

func main() {
	// Example 1: [1,0,0,1,0,1], k=2 -> 1
	nums1 := []int{1, 0, 0, 1, 0, 1}
	k1 := 2
	fmt.Printf("minMoves(%v, %d) = %d (expected 1)\n", nums1, k1, minMoves(nums1, k1))

	// Example 2: [1,0,0,0,0,0,1,1], k=3 -> 5
	nums2 := []int{1, 0, 0, 0, 0, 0, 1, 1}
	k2 := 3
	fmt.Printf("minMoves(%v, %d) = %d (expected 5)\n", nums2, k2, minMoves(nums2, k2))

	// Example 3: [1,1,0,1], k=2 -> 0
	nums3 := []int{1, 1, 0, 1}
	k3 := 2
	fmt.Printf("minMoves(%v, %d) = %d (expected 0)\n", nums3, k3, minMoves(nums3, k3))

	// Example from problem description: [1,0,0,1,0,1], k=2 -> 1 (same as example 1)
}
