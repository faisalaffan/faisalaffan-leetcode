package main

// LeetCode #2192: All Ancestors of a Node in a Directed Acyclic Graph
// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import (
	"fmt"
	"sort"
)

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for _, parent := range graph[node] {
				if !visited[parent] {
					visited[parent] = true
					queue = append(queue, parent)
				}
			}
		}
		ancestors := []int{}
		for j := 0; j < n; j++ {
			if j != i && visited[j] {
				ancestors = append(ancestors, j)
			}
		}
		sort.Ints(ancestors)
		result[i] = ancestors
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))

	// Test case 2
	fmt.Println(getAncestors(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}))
}
