package main

// LeetCode #461: Hamming Distance
// https://leetcode.com/problems/hamming-distance/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		xor &= xor - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingDistance(1, 4))
	fmt.Println(HammingDistance(3, 1))
}
