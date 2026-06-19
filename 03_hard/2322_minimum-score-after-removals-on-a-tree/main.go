package main

import (
	"fmt"
	"math"
)

// 2322. Minimum Score After Removals on a Tree
// ----------------------------------------------------------------
// Given a tree with n nodes, each having a value.  Remove exactly two edges,
// splitting the tree into three connected components.  The "score" of a
// component is the XOR of all node values in that component.
// Minimise max(score1, score2, score3) - min(score1, score2, score3).
//
// n ≤ 1000, so O(n²) is fine.
// Idea:
//   1. Root the tree at 0.
//   2. Compute subtree XOR via DFS (post‑order).
//   3. Enumerate every unordered pair of edges (i, j), i ≠ j.
//      Let subtree_xor[v] = XOR of all nodes in the subtree rooted at v.
//
//      Relationship between the two edges:
//       a) One edge is a descendant of the other.
//          If edge a (removing edge at 'a') is an ancestor of edge b,
//          components are:
//            comp1 = subtree_xor[b]
//            comp2 = subtree_xor[a] XOR subtree_xor[b]
//            comp3 = total_xor XOR subtree_xor[a]
//
//       b) The two edges are in unrelated subtrees.
//          components are:
//            comp1 = subtree_xor[a]
//            comp2 = subtree_xor[b]
//            comp3 = total_xor XOR subtree_xor[a] XOR subtree_xor[b]
//
//   Also consider removing edge a and b where one is the parent of another
//   (same as case a).

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Parent and subtree XOR via DFS.
	parent := make([]int, n)
	order := make([]int, 0, n) // DFS order
	subXor := make([]int, n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		parent[u] = p
		order = append(order, u)
		subXor[u] = nums[u]
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(v, u)
			subXor[u] ^= subXor[v]
		}
	}
	dfs(0, -1)
	total := subXor[0]

	// Pre‑compute ancestors for O(1) descendant check.
	// tin / tout using Euler tour.
	tin := make([]int, n)
	tout := make([]int, n)
	time := 0
	var euler func(u int)
	euler = func(u int) {
		tin[u] = time
		time++
		for _, v := range adj[u] {
			if v == parent[u] {
				continue
			}
			euler(v)
		}
		tout[u] = time
		time++
	}
	euler(0)

	isAncestor := func(a, b int) bool {
		// Is 'a' an ancestor of 'b'?
		return tin[a] <= tin[b] && tout[b] <= tout[a]
	}

	best := math.MaxInt32
	for i := 0; i < n; i++ {
		if i == 0 {
			continue // can't cut above root
		}
		for j := i + 1; j < n; j++ {
			if j == 0 {
				continue
			}
			var x, y, z int
			if isAncestor(i, j) {
				// i is ancestor of j
				x = subXor[j]
				y = subXor[i] ^ subXor[j]
				z = total ^ subXor[i]
			} else if isAncestor(j, i) {
				// j is ancestor of i
				x = subXor[i]
				y = subXor[j] ^ subXor[i]
				z = total ^ subXor[j]
			} else {
				x = subXor[i]
				y = subXor[j]
				z = total ^ subXor[i] ^ subXor[j]
			}
			score := max3(x, y, z) - min3(x, y, z)
			if score < best {
				best = score
			}
		}
	}
	return best
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func MinimumScoreAfterRemovalsOnATree() interface{} {
	nums := []int{1, 5, 5, 4, 11}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	return minimumScore(nums, edges)
}

func main() {
	fmt.Println(MinimumScoreAfterRemovalsOnATree())

	// Example from problem:
	// nums = [1,5,5,4,11], edges = [[0,1],[1,2],[1,3],[3,4]]
	// The minimum score after removing two edges should be 9.
	if got := minimumScore([]int{1, 5, 5, 4, 11},
		[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}); got != 9 {
		fmt.Printf("FAIL example: got %d, want 9\n", got)
	}

	// Another test.
	if got := minimumScore([]int{5, 5, 2, 4, 4, 2},
		[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}); got != 2 {
		fmt.Printf("FAIL line: got %d, want 2\n", got)
	}

	fmt.Println("Done testing 2322.")
}
