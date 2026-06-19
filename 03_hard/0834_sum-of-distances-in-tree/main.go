package main

// LeetCode #834: Sum of Distances in Tree
// https://leetcode.com/problems/sum-of-distances-in-tree/
// Difficulty: Hard
// Approach: Rerooting DP. First DFS from root to get subtree sizes and sum of distances
// from root. Second DFS to compute answers for all nodes using reroot formula:
// ans[child] = ans[parent] + n - 2*subtree[child]

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	subtree := make([]int, n)
	ans := make([]int, n)

	var dfs1 func(u, parent int)
	dfs1 = func(u, parent int) {
		subtree[u] = 1
		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			dfs1(v, u)
			subtree[u] += subtree[v]
			ans[0] += subtree[v]
		}
	}
	dfs1(0, -1)

	var dfs2 func(u, parent int)
	dfs2 = func(u, parent int) {
		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			ans[v] = ans[u] + n - 2*subtree[v]
			dfs2(v, u)
		}
	}
	dfs2(0, -1)

	return ans
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	// Expected: [8 12 6 10 10 10]

	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
	// Expected: [0]
}
