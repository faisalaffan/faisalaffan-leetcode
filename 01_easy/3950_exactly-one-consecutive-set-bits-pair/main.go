package main

// LeetCode #3950: Exactly One Consecutive Set Bits Pair
// https://leetcode.com/problems/exactly-one-consecutive-set-bits-pair/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(6))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(5))
	fmt.Println(ExactlyOneConsecutiveSetBitsPair(3))
}

// Time: O(1)
// Space: O(1)
func ExactlyOneConsecutiveSetBitsPair(n int) bool {
	m := n & (n >> 1)
	return m > 0 && m&(m-1) == 0
}
