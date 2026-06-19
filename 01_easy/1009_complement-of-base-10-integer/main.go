package main

// LeetCode #1009: Complement of Base 10 Integer
// https://leetcode.com/problems/complement-of-base-10-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5))  // 2
	fmt.Println(bitwiseComplement(7))  // 0
	fmt.Println(bitwiseComplement(10)) // 5
	fmt.Println(bitwiseComplement(0))  // 1
}

// bitwiseComplement returns the complement of a base-10 integer's binary representation.
// Time: O(log n). Space: O(1).
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	mask := n
	// Set all bits to the right of the MSB
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return ^n & mask
}
