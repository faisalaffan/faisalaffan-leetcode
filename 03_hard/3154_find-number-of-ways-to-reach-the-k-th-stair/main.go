package main

// LeetCode #3154: Find Number of Ways to Reach the K-th Stair
// https://leetcode.com/problems/find-number-of-ways-to-reach-the-k-th-stair/
// Difficulty: Hard
//
// Start at stair 1 with jump = 0.
// Operations:
//   - Go down to i-1 (cannot be used consecutively or below stair 0)
//   - Go up to i + 2^jump, then jump++
//
// Count total ways to reach stair k. May pass through k and come back.
// After `up` up-jumps, position = 2^up. Need down = 2^up - k down moves.
// Down moves must be <= up+1 (inserted into up+1 gaps, no two consecutive).
// Ways = C(up+1, down). Sum over all valid (up, down) pairs.

import (
	"fmt"
)

func comb(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func waysToReachStair(k int) int {
	ans := 0
	for up := 0; up <= 31; up++ {
		power := 1 << uint(up)
		down := power - k
		if down >= 0 && down <= up+1 {
			ans += comb(up+1, down)
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", waysToReachStair(0))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", waysToReachStair(1))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", waysToReachStair(2))
	// Expected: ?

	// Test case 4: larger
	fmt.Println("Test 4:", waysToReachStair(10))
}
