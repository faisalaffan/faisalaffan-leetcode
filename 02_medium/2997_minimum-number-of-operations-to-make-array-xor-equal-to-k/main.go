package main

// LeetCode #2997: Minimum Number of Operations to Make Array XOR Equal to K
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minOperationsXOR([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperationsXOR([]int{2, 0, 2, 0}, 0))
}

func minOperationsXOR(nums []int, k int) (ans int) {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	return bits.OnesCount(uint(xor ^ k))
}
