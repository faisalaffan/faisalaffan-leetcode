package main

// LeetCode #323: Number of Connected Components in an Undirected Graph
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Difficulty: Medium [Paid]
// Time: O(V + E) | Space: O(V)

import "fmt"

func countComponents(n int, edges [][]int) int {
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

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
			n--
		}
	}

	for _, e := range edges {
		union(e[0], e[1])
	}
	return n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
	// Expected: 1

	// Test case 3: No edges
	fmt.Println("Test 3:", countComponents(3, [][]int{}))
	// Expected: 3
}
