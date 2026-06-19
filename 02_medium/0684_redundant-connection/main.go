package main

// LeetCode #684: Redundant Connection
// https://leetcode.com/problems/redundant-connection/
// Difficulty: Medium
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		parent[px] = py
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
