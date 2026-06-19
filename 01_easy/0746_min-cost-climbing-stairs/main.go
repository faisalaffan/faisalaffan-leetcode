package main

// LeetCode #746: Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                    // 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
}

// minCostClimbingStairs returns the minimum cost to reach the top of the stairs.
// Time: O(n). Space: O(1).
func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		a, b = b, min(a, b)+cost[i]
	}
	return min(a, b)
}
