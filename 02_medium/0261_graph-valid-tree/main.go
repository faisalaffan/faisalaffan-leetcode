package main

// LeetCode #261: Graph Valid Tree
// https://leetcode.com/problems/graph-valid-tree/
// Difficulty: Medium [Paid]
// Time: O(V+E), Space: O(V+E)

import "fmt"

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
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
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
	fmt.Println(validTree(4, [][]int{{0, 1}, {2, 3}}))
}
