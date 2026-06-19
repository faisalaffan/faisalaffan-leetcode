package main

// LeetCode #1015: Smallest Integer Divisible by K
// https://leetcode.com/problems/smallest-integer-divisible-by-k/
// Difficulty: Medium
//
// Approach: Build remainder incrementally. Keep dividing.
//   N = 1, 11, 111, ... = N*10 + 1 (mod K)
//   If N % K == 0, return length. If remainder repeats, return -1.
// Time: O(K)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(smallestRepunitDivByK(1))  // 1
	fmt.Println(smallestRepunitDivByK(2))  // -1
	fmt.Println(smallestRepunitDivByK(3))  // 3
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 0
	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length
		}
	}

	return -1
}
