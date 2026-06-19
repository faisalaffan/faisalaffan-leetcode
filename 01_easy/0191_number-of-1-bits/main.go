package main

// LeetCode #191: Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func HammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingWeight(11))    // 3
	fmt.Println(HammingWeight(128))   // 1
	fmt.Println(HammingWeight(4294967293)) // 31
}
