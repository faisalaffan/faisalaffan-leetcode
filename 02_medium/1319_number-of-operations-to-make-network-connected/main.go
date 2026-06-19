package main

// LeetCode #1319: Number of Operations to Make Network Connected
// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}})) // 1

	// Test case 2
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}})) // 2

	// Test case 3
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}})) // -1
}

// Time: O(n + connections) for union-find operations
// Space: O(n) for parent and rank arrays
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1 // not enough cables
	}

	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if rank[ra] < rank[rb] {
			parent[ra] = rb
		} else if rank[ra] > rank[rb] {
			parent[rb] = ra
		} else {
			parent[rb] = ra
			rank[ra]++
		}
	}

	for _, conn := range connections {
		union(conn[0], conn[1])
	}

	components := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			components++
		}
	}

	return components - 1
}
