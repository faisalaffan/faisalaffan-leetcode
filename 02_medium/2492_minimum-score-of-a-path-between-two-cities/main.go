package main

// LeetCode #2492: Minimum Score of a Path Between Two Cities
// https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// DFS from city 1, find min edge in its connected component (must include city n).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}})) // 5
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))            // 2
}

func minScore(n int, roads [][]int) int {
	graph := make([][][2]int, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		graph[a] = append(graph[a], [2]int{b, d})
		graph[b] = append(graph[b], [2]int{a, d})
	}

	visited := make([]bool, n+1)
	ans := math.MaxInt32

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, edge := range graph[u] {
			v, d := edge[0], edge[1]
			if d < ans {
				ans = d
			}
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(1)

	return ans
}
