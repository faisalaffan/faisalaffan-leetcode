package main

// LeetCode #797: All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	path = append(path, 0)

	var dfs func(node int)
	dfs = func(node int) {
		if node == len(graph)-1 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, neighbor := range graph[node] {
			path = append(path, neighbor)
			dfs(neighbor)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}
