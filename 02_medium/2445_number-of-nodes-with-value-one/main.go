package main

// LeetCode #2445: Number of Nodes With Value One
// https://leetcode.com/problems/number-of-nodes-with-value-one/
// Difficulty: Medium
// Time: O(sqrt(n) + q * log n) | Space: O(log n)
// Complete binary tree with n nodes. Each query k toggles depths divisible by k.
// Count nodes at depths with odd toggle count.

import "fmt"

func main() {
	// Complete binary tree with 5 nodes: depths are 1(1), 2(2), 3(2 nodes since incomplete)
	// Queries: k=1 toggles all depths, k=2 toggles depths 2 only, k=3 toggles depth 3 only
	fmt.Println(numberOfNodes(5, []int{1, 2})) // 3 (depths 2,3 toggled)

	// n=6, depths: 1(1), 2(2), 3(3)
	fmt.Println(numberOfNodes(6, []int{3})) // 3
}

func numberOfNodes(n int, queries []int) int {
	// Compute depth range
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	if maxDepth == 0 {
		return 0
	}

	// Nodes at each depth (excluding last might be incomplete)
	nodesAtDepth := make([]int, maxDepth+1)
	remaining := n
	for d := 1; d <= maxDepth; d++ {
		level := 1 << (d - 1) // 2^(d-1)
		if remaining >= level {
			nodesAtDepth[d] = level
			remaining -= level
		} else {
			nodesAtDepth[d] = remaining
			remaining = 0
		}
	}

	// Toggle depths that are multiples of each query k
	toggle := make([]int, maxDepth+1)
	for _, k := range queries {
		if k <= maxDepth {
			for d := k; d <= maxDepth; d += k {
				toggle[d] ^= 1
			}
		}
	}

	ans := 0
	for d := 1; d <= maxDepth; d++ {
		if toggle[d] == 1 {
			ans += nodesAtDepth[d]
		}
	}
	return ans
}
