package main

// LeetCode #3209: Number of Subarrays With AND Value of K
// https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/
// Difficulty: Hard
//
// Count the number of subarrays whose bitwise AND equals exactly k.
// For each position, maintain a map of distinct AND values for subarrays
// ending at that position. Since AND only decreases, there are at most
// O(log MAX) distinct values per position.

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1}
	k := 1
	fmt.Println(numberOfSubarraysWithAndValueOfK(nums, k)) // expected: 6
}

func numberOfSubarraysWithAndValueOfK(nums []int, k int) int64 {
	var ans int64 = 0
	// cur map: AND value -> count for subarrays ending at current position
	cur := make(map[int]int)

	for _, x := range nums {
		nxt := make(map[int]int)
		// Start new subarray consisting of just x
		nxt[x] = 1

		// Extend existing subarrays: AND with x
		for andVal, cnt := range cur {
			key := andVal & x
			nxt[key] += cnt
		}

		// Add count for target k
		if cnt, ok := nxt[k]; ok {
			ans += int64(cnt)
		}

		cur = nxt
	}

	return ans
}
