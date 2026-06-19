package main

// LeetCode #3370: Smallest Number With All Set Bits
// https://leetcode.com/problems/smallest-number-with-all-set-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestNumberWithAllSetBits(5))
	fmt.Println(SmallestNumberWithAllSetBits(10))
	fmt.Println(SmallestNumberWithAllSetBits(3))
}

// SmallestNumberWithAllSetBits returns the smallest number >= n whose binary representation consists of all 1s.
// Time: O(log n). Space: O(1).
func SmallestNumberWithAllSetBits(n int) int {
	result := 1
	for result < n {
		result = (result << 1) | 1
	}
	return result
}
