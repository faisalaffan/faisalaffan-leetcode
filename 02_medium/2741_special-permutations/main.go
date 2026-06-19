package main

// LeetCode #2741: Special Permutations
// https://leetcode.com/problems/special-permutations/
// Difficulty: Medium
// Time: O(n^2 * 2^n) | Space: O(n * 2^n)

import "fmt"

func SpecialPermutations(nums []int) int {
	n := len(nums)
	const mod = 1_000_000_007

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for mask := 1; mask < 1<<n; mask++ {
		for last := 0; last < n; last++ {
			if dp[mask][last] == 0 {
				continue
			}
			for next := 0; next < n; next++ {
				if mask&(1<<next) != 0 {
					continue
				}
				if nums[last]%nums[next] == 0 || nums[next]%nums[last] == 0 {
					newMask := mask | (1 << next)
					dp[newMask][next] = (dp[newMask][next] + dp[mask][last]) % mod
				}
			}
		}
	}

	var result int
	fullMask := (1 << n) - 1
	for i := 0; i < n; i++ {
		result = (result + dp[fullMask][i]) % mod
	}
	return result
}

func main() {
	fmt.Println(SpecialPermutations([]int{1, 2, 3}))
	fmt.Println(SpecialPermutations([]int{2, 3, 6}))
}
