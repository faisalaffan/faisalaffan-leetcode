package main

// LeetCode #2836: Maximize Value of Function in a Ball Passing Game
// https://leetcode.com/problems/maximize-value-of-function-in-a-ball-passing-game/
// Difficulty: Hard
//
// Binary lifting. dp[i][j] = node reached after 2^j steps from i.
// sum[i][j] = sum of node IDs along the path of length 2^j from i (inclusive).
// For each starting node, compute f(i,k) by decomposing k into binary.
// O(N log K) time, O(N log K) space.

import "fmt"

func getMaxFunctionValue(receiver []int, k int64) int64 {
	n := len(receiver)
	logK := 0
	for (int64(1) << logK) <= k {
		logK++
	}

	dp := make([][]int, n)
	sum := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, logK)
		sum[i] = make([]int64, logK)
		dp[i][0] = receiver[i]
		sum[i][0] = int64(i)
	}

	for j := 1; j < logK; j++ {
		for i := 0; i < n; i++ {
			mid := dp[i][j-1]
			dp[i][j] = dp[mid][j-1]
			sum[i][j] = sum[i][j-1] + sum[mid][j-1]
		}
	}

	maxVal := int64(0)
	for start := 0; start < n; start++ {
		total := int64(0)
		cur := start
		remaining := k
		bit := 0
		for remaining > 0 {
			if remaining&1 == 1 {
				total += sum[cur][bit]
				cur = dp[cur][bit]
			}
			remaining >>= 1
			bit++
		}
		total += int64(cur)
		if total > maxVal {
			maxVal = total
		}
	}

	return maxVal
}

func main() {
	// Example: receiver=[2,0,1], k=4 => 6
	fmt.Println(getMaxFunctionValue([]int{2, 0, 1}, 4))

	// k=1
	fmt.Println(getMaxFunctionValue([]int{1, 0}, 1))

	// Larger k
	fmt.Println(getMaxFunctionValue([]int{2, 0, 1}, 10))

	// Self-loop (each node passes to itself)
	fmt.Println(getMaxFunctionValue([]int{0, 1}, 3))

	// Chain pattern
	fmt.Println(getMaxFunctionValue([]int{1, 2, 0}, 2))

	// All nodes point to themselves
	fmt.Println(getMaxFunctionValue([]int{0, 1, 2, 3}, 5))

	// Cycle of length 3
	fmt.Println(getMaxFunctionValue([]int{1, 2, 0}, 7))

	// Single node pointing to itself
	fmt.Println(getMaxFunctionValue([]int{0}, 10))

	// Two-node cycle
	fmt.Println(getMaxFunctionValue([]int{1, 0}, 6))

	// k=0 (just the starting node, k >= 1 per constraints but testing)
	// Note: k=0 isn't in constraints (1 <= k <= 1e10), but let's test
	// fmt.Println(getMaxFunctionValue([]int{0, 1, 2}, 0))

	// Large k with small cycle
	fmt.Println(getMaxFunctionValue([]int{1, 2, 0}, 100))

	// Linear chain (no cycle, all point to next, last self-loop)
	fmt.Println(getMaxFunctionValue([]int{1, 2, 3, 3}, 4))
}
