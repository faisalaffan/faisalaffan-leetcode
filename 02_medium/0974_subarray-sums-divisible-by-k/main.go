package main

// LeetCode #974: Subarray Sums Divisible by K
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Difficulty: Medium
//
// Approach: Prefix sum + modulo counting
// Time: O(n)
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{5}, 9))                  // 0
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))           // 2
}

func subarraysDivByK(nums []int, k int) int {
	modCount := make(map[int]int)
	modCount[0] = 1
	prefixSum := 0
	result := 0

	for _, n := range nums {
		prefixSum += n
		mod := prefixSum % k
		if mod < 0 {
			mod += k
		}
		result += modCount[mod]
		modCount[mod]++
	}

	return result
}
