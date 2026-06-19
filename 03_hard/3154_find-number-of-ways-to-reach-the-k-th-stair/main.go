package main

// LeetCode #3154: Find Number of Ways to Reach the K-th Stair
// https://leetcode.com/problems/find-number-of-ways-to-reach-the-k-th-stair/
// Difficulty: Hard
//
// You start at stair 1 with jump = 0. Operations:
//   1. Go down to i-1 (cannot be used consecutively or on stair 0).
//   2. Go up to i + 2^jump, then jump++.
// Count total ways to reach stair k. You may pass through k and come back.
//
// Mathematical approach: after `up` up-jumps, position = 2^up.
// Need `down = 2^up - k` down moves (cannot be consecutive).
// Down moves must be <= up+1 (inserted into up+1 gaps).
// Ways = C(up+1, down). Sum over all valid (up, down) pairs.

import (
	"fmt"
)

func main() {
	fmt.Println(waysToReachStair(0)) // expected: 2
	fmt.Println(waysToReachStair(1)) // expected: 4
}

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
	// up-jumps can be from 0 to 30 (2^30 > 10^9, k <= 10^9)
	for up := 0; up <= 31; up++ {
		down := (1 << uint(up)) - k
		if down >= 0 && down <= up+1 {
			ans += comb(up+1, down)
		}
	}
	return ans
}
