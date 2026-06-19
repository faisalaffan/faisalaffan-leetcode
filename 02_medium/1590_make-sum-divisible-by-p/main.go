package main

// LeetCode #1590: Make Sum Divisible by P
// https://leetcode.com/problems/make-sum-divisible-by-p/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(MinSubarray([]int{6, 3, 5, 2}, 9))
	fmt.Println(MinSubarray([]int{1, 2, 3}, 7))
}

func MinSubarray(nums []int, p int) int {
	// Time: O(N), Space: O(N)
	n := len(nums)

	// Total sum modulo p
	totalSum := 0
	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum // the remainder we need to remove
	if target == 0 {
		return 0
	}

	// Map from prefix sum modulo p to index
	prefixMap := make(map[int]int)
	prefixMap[0] = -1
	prefixSum := 0
	minLen := n

	for i, num := range nums {
		prefixSum = (prefixSum + num) % p
		// We need prefixSum - prefixSum[j] ≡ target (mod p)
		// => prefixSum[j] ≡ prefixSum - target (mod p)
		needed := (prefixSum - target + p) % p
		if j, exists := prefixMap[needed]; exists {
			if i-j < minLen {
				minLen = i - j
			}
		}
		prefixMap[prefixSum] = i
	}

	if minLen == n {
		return -1
	}
	return minLen
}
