package main

// LeetCode #2786: Visit Array Positions to Maximize Score
// https://leetcode.com/problems/visit-array-positions-to-maximize-score/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func VisitArrayPositionsToMaximizeScore(nums []int, x int) int64 {
	n := len(nums)
	// dp[0] = max score ending at even parity, dp[1] = max score ending at odd parity
	dp := [2]int64{int64(nums[0]), int64(nums[0])}
	parity0 := nums[0] & 1

	other := 1 - parity0
	dp[other] = int64(nums[0]) - int64(x)
	if dp[other] < int64(nums[0]) {
		dp[other] = int64(nums[0])
	}

	best := int64(nums[0])
	for i := 1; i < n; i++ {
		p := nums[i] & 1
		// Option 1: Start here
		cur := int64(nums[i])
		// Option 2: Extend from prev with same parity
		if int64(nums[i])+dp[p] > cur {
			cur = int64(nums[i]) + dp[p]
		}
		// Option 3: Extend from prev with different parity (pay x)
		otherP := 1 - p
		if int64(nums[i])-int64(x)+dp[otherP] > cur {
			cur = int64(nums[i]) - int64(x) + dp[otherP]
		}
		if cur > dp[p] {
			dp[p] = cur
		}
		if cur > best {
			best = cur
		}
	}

	return best
}

func main() {
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 3, 6, 1, 9, 2}, 5))
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 4, 6, 8}, 3))
}
