package main

// LeetCode #2425: Bitwise XOR of All Pairings
// https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)
// XOR of all pairings = if len(nums2) odd, XOR all nums1; if len(nums1) odd, XOR all nums2.

import "fmt"

func main() {
	fmt.Println(xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0})) // 13
	fmt.Println(xorAllNums([]int{1, 2}, []int{3, 4}))            // 0
}

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			ans ^= v
		}
	}
	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			ans ^= v
		}
	}
	return ans
}
