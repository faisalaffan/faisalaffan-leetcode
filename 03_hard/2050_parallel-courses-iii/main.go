package main

// LeetCode #2050: Parallel Courses III
// https://leetcode.com/problems/parallel-courses-iii/
// Difficulty: Hard
// Approach: Topological Sort + DP

import "fmt"

func minimumTime(n int, relations [][]int, time []int) int {
	// Build graph and indegree
	graph := make([][]int, n+1)
	indeg := make([]int, n+1)
	for _, r := range relations {
		prev, next := r[0], r[1]
		graph[prev] = append(graph[prev], next)
		indeg[next]++
	}

	// dp[i] = earliest completion time for course i (1-indexed)
	dp := make([]int, n+1)
	queue := make([]int, 0)

	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			dp[i] = time[i-1]
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			// Can start v only after u finishes
			if dp[u] > dp[v] {
				dp[v] = dp[u]
			}
			indeg[v]--
			if indeg[v] == 0 {
				dp[v] += time[v-1]
				queue = append(queue, v)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println("2050. Parallel Courses III")

	// Example 1
	n1 := 3
	relations1 := [][]int{{1, 3}, {2, 3}}
	time1 := []int{3, 2, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 8)\n",
		n1, relations1, time1, minimumTime(n1, relations1, time1))

	// Example 2
	n2 := 5
	relations2 := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 12)\n",
		n2, relations2, time2, minimumTime(n2, relations2, time2))
}
