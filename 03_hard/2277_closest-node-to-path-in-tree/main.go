package main

import (
	"fmt"
)

// 2277. Closest Node to Path in Tree
// ----------------------------------------------------------------
// For each query (u, v, node), return the node on the path between u and v
// that is closest to the query node.
//
// Solution: Binary Lifting for LCA.
//   dist(a, b) = depth[a] + depth[b] - 2*depth[LCA(a,b)]
//   The closest node on path(u,v) to x is the *deepest* among:
//     LCA(u,v), LCA(x,u), LCA(x,v)  — provided it lies on path(u,v).

func closestNode(n int, edges [][]int, query [][]int) []int {
	// Build adjacency.
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Binary lifting pre‑computation.
	LOG := 0
	for (1 << LOG) <= n {
		LOG++
	}
	up := make([][]int, n)
	depth := make([]int, n)
	for i := range up {
		up[i] = make([]int, LOG)
	}

	var dfs func(v, p int)
	dfs = func(v, p int) {
		up[v][0] = p
		for k := 1; k < LOG; k++ {
			up[v][k] = up[up[v][k-1]][k-1]
		}
		for _, to := range adj[v] {
			if to == p {
				continue
			}
			depth[to] = depth[v] + 1
			dfs(to, v)
		}
	}
	dfs(0, 0) // tree is connected, root at 0

	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}
		// Lift a up to depth[b].
		diff := depth[a] - depth[b]
		for k := 0; k < LOG; k++ {
			if diff&(1<<k) != 0 {
				a = up[a][k]
			}
		}
		if a == b {
			return a
		}
		for k := LOG - 1; k >= 0; k-- {
			if up[a][k] != up[b][k] {
				a = up[a][k]
				b = up[b][k]
			}
		}
		return up[a][0]
	}

	dist := func(a, b int) int {
		return depth[a] + depth[b] - 2*depth[lca(a, b)]
	}

	// Check whether node p lies on the path between u and v.
	onPath := func(u, v, p int) bool {
		return dist(u, p)+dist(p, v) == dist(u, v)
	}

	ans := make([]int, len(query))
	for i, q := range query {
		u, v, x := q[0], q[1], q[2]
		lcauv := lca(u, v)
		candidates := []int{lcauv, lca(x, u), lca(x, v)}
		best, bestDepth := -1, -1
		for _, c := range candidates {
			if onPath(u, v, c) && depth[c] > bestDepth {
				best = c
				bestDepth = depth[c]
			}
		}
		ans[i] = best
	}
	return ans
}

// ---------------------------------------------------------------------------
//  LeetCode-style wrapper

func ClosestNodeToPathInTree() interface{} {
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	queries := [][]int{{1, 3, 0}, {4, 2, 0}}
	return closestNode(5, edges, queries)
}

func main() {
	fmt.Println(ClosestNodeToPathInTree())

	// ---- simple test ----
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}
	got := closestNode(5, edges, [][]int{{1, 3, 0}})
	fmt.Println("closestNode(1,3,0) =", got[0], "(expected 1)")
	got = closestNode(5, edges, [][]int{{4, 2, 0}})
	fmt.Println("closestNode(4,2,0) =", got[0], "(expected 2)")
	fmt.Println("Done testing 2277.")
}
