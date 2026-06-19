package main

// LeetCode #1029: Two City Scheduling
// https://leetcode.com/problems/two-city-scheduling/
// Difficulty: Medium
//
// Approach: Sort by difference (costA - costB), send first half to A, rest to B
// Time: O(n log n)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}})) // 110
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}})) // 1859
}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	n := len(costs) / 2
	total := 0
	for i := 0; i < n; i++ {
		total += costs[i][0]
	}
	for i := n; i < 2*n; i++ {
		total += costs[i][1]
	}
	return total
}
