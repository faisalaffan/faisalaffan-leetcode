package main

// LeetCode #454: 4Sum II
// https://leetcode.com/problems/4sum-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			sumMap[a+b]++
		}
	}

	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			count += sumMap[-(c + d)]
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
	// Expected: 6
}
