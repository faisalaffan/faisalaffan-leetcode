package main

// LeetCode #644: Maximum Average Subarray II
// https://leetcode.com/problems/maximum-average-subarray-ii/
// Difficulty: Hard [Paid]

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
		{[]int{1, 2, 3, 4, 5}, 2, 4.5},
		{[]int{-1, -2, -3, -4, -5}, 2, -1.5},
		{[]int{0, 4, 0, 3, 2}, 1, 4.0},
		{[]int{10, 20, 30, 40, 50}, 5, 30.0},
		{[]int{1, 12, -5, -6, 50, 3}, 3, 15.666666666666666},
		{[]int{0, 0, 0, 0}, 2, 0.0},
	}

	epsilon := 1e-5
	for _, tc := range testCases {
		got := findMaxAverage(tc.nums, tc.k)
		diff := got - tc.want
		if diff < 0 {
			diff = -diff
		}
		status := "PASS"
		if diff > epsilon {
			status = "FAIL"
		}
		fmt.Printf("%s: findMaxAverage(%v, %d) = %v (want %v)\n", status, tc.nums, tc.k, got, tc.want)
	}
}

func findMaxAverage(nums []int, k int) float64 {
	// Binary search on the average value
	left, right := -10000.0, 10000.0

	for right-left > 1e-6 {
		mid := left + (right-left)/2
		if canAchieve(nums, k, mid) {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

func canAchieve(nums []int, k int, target float64) bool {
	n := len(nums)
	// prefix[i] = sum (nums[j] - target) for j = 0..i-1
	prefix := make([]float64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + float64(nums[i]) - target
	}

	// Check if there's a subarray of length >= k with sum >= 0
	minPrefix := 0.0
	for i := k; i <= n; i++ {
		if prefix[i]-minPrefix >= 0 {
			return true
		}
		// Update minPrefix for next iteration
		if prefix[i-k+1] < minPrefix {
			minPrefix = prefix[i-k+1]
		}
	}
	return false
}
