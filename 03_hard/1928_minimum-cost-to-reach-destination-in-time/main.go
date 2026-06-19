package main

// LeetCode #1928: Minimum Cost to Reach Destination in Time
// https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/
// Difficulty: Hard - Dijkstra (DP on time)
// State: dp[city][time] = min cost to reach city at exactly that time.
// We process cities in increasing time order.

import (
	"fmt"
	"math"
)

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	// dp[t][v] = min cost to reach node v at time t
	dp := make([][]int, maxTime+1)
	for t := range dp {
		dp[t] = make([]int, n)
		for v := range dp[t] {
			dp[t][v] = math.MaxInt32
		}
	}
	dp[0][0] = passingFees[0]

	// Build adjacency list
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, t})
		adj[v] = append(adj[v], [2]int{u, t})
	}

	ans := math.MaxInt32
	for t := 0; t <= maxTime; t++ {
		for v := 0; v < n; v++ {
			if dp[t][v] == math.MaxInt32 {
				continue
			}
			if v == n-1 {
				if dp[t][v] < ans {
					ans = dp[t][v]
				}
			}
			for _, edge := range adj[v] {
				to, w := edge[0], edge[1]
				nt := t + w
				if nt <= maxTime {
					cost := dp[t][v] + passingFees[to]
					if cost < dp[nt][to] {
						dp[nt][to] = cost
					}
				}
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	// Example 1
	maxTime := 30
	edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime, edges, passingFees)) // Expected: 11

	// Example 2
	maxTime2 := 29
	edges2 := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees2 := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime2, edges2, passingFees2)) // Expected: 48 (take path 0->3->4->5)

	// Example 3
	maxTime3 := 1
	edges3 := [][]int{{0, 1, 1}, {1, 2, 1}}
	passingFees3 := []int{0, 1, 2}
	fmt.Println(minCost(maxTime3, edges3, passingFees3)) // Expected: -1
}
