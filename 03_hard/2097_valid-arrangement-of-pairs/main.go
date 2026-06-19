package main

// LeetCode #2097: Valid Arrangement of Pairs
// https://leetcode.com/problems/valid-arrangement-of-pairs/
// Difficulty: Hard
//
// Approach: Eulerian Path (Hierholzer's Algorithm).
// Build adjacency list, compute in-degree and out-degree.
// Start node = node with out-degree > in-degree, or any node with outgoing edges.
// Use iterative DFS to reconstruct the path, then reverse.

import (
	"fmt"
)

func main() {
	// Example from problem statement
	pairs1 := [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}
	fmt.Printf("validArrangement(%v) = %v (expected [[11 9] [9 4] [4 5] [5 1]])\n",
		pairs1, validArrangement(pairs1))

	// Additional tests
	pairs2 := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Printf("validArrangement(%v) = %v\n", pairs2, validArrangement(pairs2))

	pairs3 := [][]int{{1, 2}, {2, 1}}
	fmt.Printf("validArrangement(%v) = %v\n", pairs3, validArrangement(pairs3))
}

func validArrangement(pairs [][]int) [][]int {
	adj := make(map[int][]int)
	inDeg := make(map[int]int)
	outDeg := make(map[int]int)

	for _, p := range pairs {
		u, v := p[0], p[1]
		adj[u] = append(adj[u], v)
		outDeg[u]++
		inDeg[v]++
	}

	// Find start node: node with out > in, or any node with edges
	start := pairs[0][0]
	for node := range adj {
		if outDeg[node] > inDeg[node] {
			start = node
			break
		}
	}

	// Hierholzer's algorithm (iterative)
	stack := []int{start}
	path := []int{}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(adj[cur]) > 0 {
			next := adj[cur][0]
			adj[cur] = adj[cur][1:]
			stack = append(stack, next)
		} else {
			path = append(path, cur)
			stack = stack[:len(stack)-1]
		}
	}

	// Reverse path to get correct order
	result := make([][]int, len(path)-1)
	for i := 0; i < len(path)-1; i++ {
		result[i] = []int{path[len(path)-1-i], path[len(path)-2-i]}
	}
	return result
}
