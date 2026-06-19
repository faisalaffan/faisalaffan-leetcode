package main

// LeetCode #3313: Find the Last Marked Nodes in Tree
// https://leetcode.com/problems/find-the-last-marked-nodes-in-tree/
// Difficulty: Hard
//
// Tree diameter approach:
// 1. DFS from node 0 to find one diameter endpoint A.
// 2. DFS from A to find the other endpoint B and distances distA.
// 3. DFS from B to get distances distB.
// 4. For each node i, answer[i] = A if distA[i] > distB[i] else B.

import "fmt"

func main() {
	// Example: n=5, edges=[[0,1],[0,2],[2,3],[2,4]] -> expected [1,3,2,3,3]
	fmt.Println(lastMarkedNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}))

	// Single edge (n=2)
	fmt.Println(lastMarkedNodes([][]int{{0, 1}}))

	// Star tree (n=4)
	fmt.Println(lastMarkedNodes([][]int{{0, 1}, {0, 2}, {0, 3}}))

	// Path graph (n=4)
	fmt.Println(lastMarkedNodes([][]int{{0, 1}, {1, 2}, {2, 3}}))

	// More complex tree
	fmt.Println(lastMarkedNodes([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
}

func lastMarkedNodes(edges [][]int) []int {
	n := len(edges) + 1
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// DFS to compute distances from a source
	var dfs func(u, parent int, dist []int)
	dfs = func(u, parent int, dist []int) {
		for _, v := range adj[u] {
			if v != parent {
				dist[v] = dist[u] + 1
				dfs(v, u, dist)
			}
		}
	}

	// Step 1: Find A (farthest from node 0)
	dist0 := make([]int, n)
	dfs(0, -1, dist0)
	a := 0
	for i := 1; i < n; i++ {
		if dist0[i] > dist0[a] {
			a = i
		}
	}

	// Step 2: Find B (farthest from A) and compute distances from A
	distA := make([]int, n)
	dfs(a, -1, distA)
	b := 0
	for i := 1; i < n; i++ {
		if distA[i] > distA[b] {
			b = i
		}
	}

	// Step 3: Compute distances from B
	distB := make([]int, n)
	dfs(b, -1, distB)

	// Step 4: For each node, answer = farther endpoint
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if distA[i] > distB[i] {
			ans[i] = a
		} else {
			ans[i] = b
		}
	}
	return ans
}
