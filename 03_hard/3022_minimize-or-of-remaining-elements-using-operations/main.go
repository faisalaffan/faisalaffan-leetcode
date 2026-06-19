package main

// LeetCode #3022: Minimize OR of Remaining Elements Using Operations
// https://leetcode.com/problems/minimize-or-of-remaining-elements-using-operations/
// Difficulty: Hard
//
// Approach: Greedy bit-by-bit
// For each bit from high to low (29..0), try to make it 0 in the final result.
// Use AND operations across subarrays to eliminate the bit. If we can split
// into <= k subarrays (using k-1 operations) such that each subarray's AND
// has no forbidden bits (tracked by mask), we can clear this bit.
// Otherwise, this bit must be 1 in the answer.

import "fmt"

func minOrAfterOperations(nums []int, k int) int {
	ans := 0
	mask := 0
	for b := 29; b >= 0; b-- {
		mask |= 1 << b
		cnt := 0
		and := -1
		for _, x := range nums {
			and &= x & mask
			if and != 0 {
				cnt++
			} else {
				and = -1
			}
		}
		if cnt > k {
			ans |= 1 << b
			mask ^= 1 << b
		}
	}
	return ans
}

func main() {
	// Example: [3,5,3,2,7], k=2 -> 3
	fmt.Println(minOrAfterOperations([]int{3, 5, 3, 2, 7}, 2))
	// Example: [7,3,15,14,2,8], k=4 -> 2
	fmt.Println(minOrAfterOperations([]int{7, 3, 15, 14, 2, 8}, 4))
}
