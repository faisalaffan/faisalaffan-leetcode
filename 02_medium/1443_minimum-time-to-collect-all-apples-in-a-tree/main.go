package main

// LeetCode #1443: Minimum Time to Collect All Apples in a Tree
// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, true, true, false})) // 8

	// Test case 2
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, false, true, false})) // 6

	// Test case 3
	fmt.Println(minTime(4, [][]int{{0, 2}, {0, 3}, {1, 2}},
		[]bool{false, true, false, false})) // 4
}

// Time: O(n) where n = number of nodes
// Space: O(n) for adjacency list and recursion stack
func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(node int) int {
		visited[node] = true
		time := 0

		for _, child := range adj[node] {
			if !visited[child] {
				childTime := dfs(child)
				if childTime > 0 || hasApple[child] {
					time += childTime + 2 // 2 for going down and back up
				}
			}
		}

		return time
	}

	return dfs(0)
}
