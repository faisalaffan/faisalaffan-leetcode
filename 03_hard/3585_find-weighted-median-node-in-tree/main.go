package main

// LeetCode #3585: Find Weighted Median Node in Tree
// https://leetcode.com/problems/find-weighted-median-node-in-tree/
// Difficulty: Hard
//
// Given a weighted tree and queries, for each query (u, v) find the weighted
// median node on the path from u to v. The weighted median is the first node x
// such that the sum of edge weights from u to x >= half of total path weight.
//
// Approach: Binary lifting (LCA) to find ancestor at certain distance.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(findMedian(4, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 1}}, [][]int{{0, 3}, {1, 3}}))
	// Example 2: single edge
	fmt.Println(findMedian(2, [][]int{{0, 1, 5}}, [][]int{{0, 1}}))
	// Edge: single node
	fmt.Println(findMedian(1, [][]int{}, [][]int{{0, 0}}))
}

const LOG = 17

func findMedian(n int, edges [][]int, queries [][]int) []int {
	if n == 0 {
		return []int{}
	}

	// Build adjacency
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	// Binary lifting tables
	up := make([][LOG]int, n)
	depth := make([]int, n)
	dist := make([]int64, n) // distance from root

	var dfs func(u, p int)
	dfs = func(u, p int) {
		up[u][0] = p
		for j := 1; j < LOG; j++ {
			if up[u][j-1] != -1 {
				up[u][j] = up[up[u][j-1]][j-1]
			} else {
				up[u][j] = -1
			}
		}
		for _, e := range adj[u] {
			if e.to == p {
				continue
			}
			depth[e.to] = depth[u] + 1
			dist[e.to] = dist[u] + int64(e.w)
			dfs(e.to, u)
		}
	}

	// Initialize up with -1
	for i := 0; i < n; i++ {
		for j := 0; j < LOG; j++ {
			up[i][j] = -1
		}
	}
	dfs(0, -1)

	// LCA
	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		// Lift u to match v's depth
		diff := depth[u] - depth[v]
		for j := 0; j < LOG; j++ {
			if diff&(1<<j) != 0 {
				u = up[u][j]
			}
		}
		if u == v {
			return u
		}
		for j := LOG - 1; j >= 0; j-- {
			if up[u][j] != up[v][j] {
				u = up[u][j]
				v = up[v][j]
			}
		}
		return up[u][0]
	}

	// Kth ancestor
	kthAncestor := func(u, k int) int {
		for j := 0; j < LOG; j++ {
			if k&(1<<j) != 0 {
				u = up[u][j]
				if u == -1 {
					return -1
				}
			}
		}
		return u
	}

	result := make([]int, len(queries))

	for qi, q := range queries {
		u, v := q[0], q[1]
		if u == v {
			result[qi] = u
			continue
		}

		w := lca(u, v)
		totalDist := dist[u] + dist[v] - 2*dist[w]
		half := (totalDist + 1) / 2 // ceil half

		// Distance from u to w
		distUtoW := dist[u] - dist[w]
		if distUtoW >= half {
			// Median is on u-w path, need to go up from u by k steps where
			// cumulative weight >= half
			// Binary lift from u to find the node
			remaining := half
			node := u
			for j := LOG - 1; j >= 0; j-- {
				if up[node][j] != -1 {
					anc := up[node][j]
					// Distance from u to anc = dist[u] - dist[anc]
					edgeDist := dist[u] - dist[anc]
					if edgeDist < remaining {
						remaining -= edgeDist
						node = anc
					}
				}
			}
			// Now remaining > 0 means the median is on edge from node to parent
			// The problem asks for the first node where cumulative >= half
			// So if the median is on an edge, we return the child (node)
			result[qi] = node
		} else {
			// Median is on v-w path
			remaining := totalDist - half // distance from v going up
			node := v
			for j := LOG - 1; j >= 0; j-- {
				if up[node][j] != -1 {
					anc := up[node][j]
					edgeDist := dist[v] - dist[anc]
					if edgeDist < remaining {
						remaining -= edgeDist
						node = anc
					}
				}
			}
			result[qi] = node
		}
	}

	return result
}
