package main

// LeetCode #2297: Jump Game VIII
// https://leetcode.com/problems/jump-game-viii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func minCost(nums []int, costs []int) int64 {
	n := len(nums)
	dp := make([]int64, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	stackGE := []int{0} // monotonic stack for >=
	stackL := []int{0}  // monotonic stack for <

	for i := 1; i < n; i++ {
		// Jump from where nums[j] <= nums[i] (looking for >=)
		for len(stackGE) > 0 && nums[stackGE[len(stackGE)-1]] <= nums[i] {
			j := stackGE[len(stackGE)-1]
			stackGE = stackGE[:len(stackGE)-1]
			if dp[j] != -1 {
				cost := dp[j] + int64(costs[i])
				if dp[i] == -1 || cost < dp[i] {
					dp[i] = cost
				}
			}
		}
		// Jump from where nums[j] > nums[i] (looking for <)
		for len(stackL) > 0 && nums[stackL[len(stackL)-1]] > nums[i] {
			j := stackL[len(stackL)-1]
			stackL = stackL[:len(stackL)-1]
			if dp[j] != -1 {
				cost := dp[j] + int64(costs[i])
				if dp[i] == -1 || cost < dp[i] {
					dp[i] = cost
				}
			}
		}
		stackGE = append(stackGE, i)
		stackL = append(stackL, i)
	}
	return dp[n-1]
}

func main() {
	// Test case 1
	fmt.Println(minCost([]int{3, 2, 4, 4, 1}, []int{3, 7, 6, 4, 2}))
	// Expected: 8

	// Test case 2
	fmt.Println(minCost([]int{0, 1, 2}, []int{1, 1, 1}))
	// Expected: 2
}
