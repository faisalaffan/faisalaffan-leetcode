package main

// LeetCode #1734: Decode XORed Permutation
// https://leetcode.com/problems/decode-xored-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded) + 1

	// XOR of all numbers from 1 to n
	totalXor := 0
	for i := 1; i <= n; i++ {
		totalXor ^= i
	}

	// XOR of encoded[1], encoded[3], encoded[5], ...
	xorOdd := 0
	for i := 1; i < len(encoded); i += 2 {
		xorOdd ^= encoded[i]
	}

	// First element = totalXor ^ xorOdd
	perm := make([]int, n)
	perm[0] = totalXor ^ xorOdd

	// Decode the rest
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}

func main() {
	fmt.Println(decode([]int{3, 1}))          // Expected: [1, 2, 3]
	fmt.Println(decode([]int{6, 5, 4, 6}))    // Expected: [2, 4, 1, 5, 3]
	fmt.Println(decode([]int{5, 6, 1, 6, 2, 1})) // Expected: [1 4 2 3 5 7 6]
}
