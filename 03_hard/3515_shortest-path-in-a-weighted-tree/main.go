package main

// LeetCode #3515: Shortest Path in a Weighted Tree
// https://leetcode.com/problems/shortest-path-in-a-weighted-tree/
// Difficulty: Hard
//
// Euler tour flattening + Fenwick tree for range updates (edge weight changes)
// and point queries (root-to-node distance).

import "fmt"

func main() {
	fmt.Println(ShortestPathInAWeightedTree(5, [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 4}, {0, 4, 5}},
		[][]int{{0, 2}, {1, 1, 2, 5}, {0, 2}}))
}

type Fenwick struct {
	tree []int64
	n    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{tree: make([]int64, n+2), n: n}
}

func (f *Fenwick) Add(idx int, val int64) {
	idx++
	for idx <= f.n+1 {
		f.tree[idx] += val
		idx += idx & -idx
	}
}

func (f *Fenwick) Sum(idx int) int64 {
	idx++
	res := int64(0)
	for idx > 0 {
		res += f.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (f *Fenwick) RangeAdd(l, r int, val int64) {
	f.Add(l, val)
	f.Add(r+1, -val)
}

func ShortestPathInAWeightedTree(n int, edges [][]int, queries [][]int) []int64 {
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	tin := make([]int, n)
	tout := make([]int, n)
	euler := make([]int, 0, 2*n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		tin[u] = len(euler)
		euler = append(euler, u)
		for _, nb := range adj[u] {
			v := nb[0]
			if v == p {
				continue
			}
			dfs(v, u)
		}
		tout[u] = len(euler) - 1
	}
	dfs(0, -1)

	ft := NewFenwick(n)
	// Initialize with edge weights on Euler tour
	// Root (0) has distance 0. For each edge (u,v,w), add w to subtree of v.
	// So query(tin[v]) gives distance from root to v.
	var init func(u, p int) int64
	init = func(u, p int) int64 {
		dist := int64(0)
		if p == -1 {
			dist = 0
		} else {
			// Find weight of edge (p,u)
			for _, nb := range adj[u] {
				if nb[0] == p {
					dist = int64(nb[1])
					break
				}
			}
			ft.RangeAdd(tin[u], tout[u], dist)
		}
		for _, nb := range adj[u] {
			v := nb[0]
			if v != p {
				init(v, u)
			}
		}
		return dist
	}
	init(0, -1)

	var result []int64
	for _, q := range queries {
		if q[0] == 0 {
			// Query distance from root to node q[1]
			result = append(result, ft.Sum(tin[q[1]]))
		} else {
			// Update: q[1] = u, q[2] = v, q[3] = new weight
			u, v, w := q[1], q[2], q[3]
			// Determine which one is child (deeper in Euler tour)
			child := u
			if tin[v] > tin[u] {
				child = v
			}
			// Actually: ft.Sum(tin[child]) is current distance. We want to set it.
			// The simplest: range update child's subtree with diff.

			// Find current distance of child
			curDist := ft.Sum(tin[child])
			// Find distance of parent
			parentNode := parentOf(adj, child, u, v)
			parentDist := int64(0)
			if parentNode != -1 {
				parentDist = ft.Sum(tin[parentNode])
			}
			newDist := parentDist + int64(w)
			adjustment := newDist - curDist
			ft.RangeAdd(tin[child], tout[child], adjustment)
		}
	}
	return result
}

func parentOf(adj [][][2]int, child, u, v int) int {
	// One of u,v is the child, the other is the parent (or could be either way)
	if u == child {
		return v
	}
	if v == child {
		return u
	}
	return -1
}
