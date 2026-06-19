package main

// LeetCode #2172: Maximum AND Sum of Array
// https://leetcode.com/problems/maximum-and-sum-of-array/
// Difficulty: Hard
//
// DP with bitmask. Double each slot (capacity 2 -> 2 slots of capacity 1).
// dp[mask] = max AND sum for the given assignment mask.

import "fmt"

func main() {
	fmt.Println(maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)) // 9
	fmt.Println(maximumANDSum([]int{1, 3, 10, 4, 7, 1}, 3)) // 16
	fmt.Println(maximumANDSum([]int{1, 2, 3}, 2))            // 4
	fmt.Println(maximumANDSum([]int{1, 2}, 1))               // 1
}

func maximumANDSum(nums []int, numSlots int) int {
	n := len(nums)
	m := 2 * numSlots // doubled slots
	total := 1 << m

	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = -1
	}

	for mask := 0; mask < total; mask++ {
		if dp[mask] < 0 {
			continue
		}
		idx := popcount(mask)
		if idx >= n {
			continue
		}
		for slot := 0; slot < m; slot++ {
			if mask&(1<<slot) == 0 {
				nm := mask | (1 << slot)
				val := dp[mask] + (nums[idx] & (slot/2 + 1))
				if val > dp[nm] {
					dp[nm] = val
				}
			}
		}
	}
	return dp[total-1]
}

func popcount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func MaximumAndSumOfArray() any {
	return maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)
}
