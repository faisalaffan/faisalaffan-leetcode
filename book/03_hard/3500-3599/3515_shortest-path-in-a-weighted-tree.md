# 3515 — Shortest Path In A Weighted Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func NewFenwick(n int) *Fenwick`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Fenwick Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3515: Shortest Path in a Weighted Tree
// https://leetcode.com/problems/shortest-path-in-a-weighted-tree/
// Difficulty: Hard
//
// Binary lifting for LCA + Fenwick tree for Euler tour range updates.
// Edge weight update: update subtree of the child node.
// Distance query: dist(root,u) + dist(root,v) - 2*dist(root,lca(u,v)).
// Edge update: find which endpoint is child, update its subtree weight delta.

import "fmt"

type Fenwick struct {
	tree []int64
	n    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{tree: make([]int64, n+2), n: n}
}

func (f *Fenwick) add(idx int, val int64) {
	idx++
	for idx <= f.n+1 {
		f.tree[idx] += val
		idx += idx & -idx
	}
}

func (f *Fenwick) sum(idx int) int64 {
	idx++
	res := int64(0)
	for idx > 0 {
		res += f.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (f *Fenwick) rangeAdd(l, r int, val int64) {
	f.add(l, val)
	f.add(r+1, -val)
}

func shortestPathWeightedTree(n int, edges [][]int, queries [][]int) []int64 {
	LOG := 17
	for 1<<LOG <= n {
		LOG++
	}

  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

  // Alokasi slice
	tin := make([]int, n)
  // Alokasi slice
	tout := make([]int, n)
  // Alokasi slice
	depth := make([]int, n)
  // Matriks 2D
	up := make([][]int, n)
  // Range loop
	for i := range up {
		up[i] = make([]int, LOG)
	}
	// edgeParent[child] = parent node (for edge weight lookup)
  // Alokasi slice
	edgeParent := make([]int, n)
  // Range loop
	for i := range edgeParent {
		edgeParent[i] = -1
	}
	// edgeWeightToParent[child] = weight of edge to parent
  // Alokasi slice
	edgeW := make([]int64, n)

  // Alokasi slice
	euler := make([]int, 0, 2*n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		tin[u] = len(euler)
		euler = append(euler, u)
		up[u][0] = p
		if p == -1 {
			up[u][0] = u
		}
		for k := 1; k < LOG; k++ {
			up[u][k] = up[up[u][k-1]][k-1]
		}
		for _, nb := range adj[u] {
			v, w := nb[0], nb[1]
			if v == p {
				continue
			}
			depth[v] = depth[u] + 1
			edgeParent[v] = u
			edgeW[v] = int64(w)
			dfs(v, u)
		}
		tout[u] = len(euler) - 1
	}
	dfs(0, 0)

	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		// Lift u to same depth
		diff := depth[u] - depth[v]
		for k := 0; k < LOG; k++ {
			if diff>>k&1 == 1 {
				u = up[u][k]
			}
		}
		if u == v {
			return u
		}
		for k := LOG - 1; k >= 0; k-- {
			if up[u][k] != up[v][k] {
				u = up[u][k]
				v = up[v][k]
			}
		}
		return up[u][0]
	}

	// Fenwick tree over Euler tour for distance updates
	ft := NewFenwick(n)

	// Initialize distances
	var initDist func(u int)
	initDist = func(u int) {
		if u == 0 {
			// distance from root to root = 0
			ft.rangeAdd(tin[u], tout[u], 0)
		} else {
			parentDist := ft.sum(tin[edgeParent[u]])
			ft.rangeAdd(tin[u], tout[u], parentDist+edgeW[u])
		}
		for _, nb := range adj[u] {
			v := nb[0]
			if v == edgeParent[u] {
				continue
			}
			initDist(v)
		}
	}
	initDist(0)

	// After initDist, each node u has dist from root = ft.sum(tin[u])
	// But the initialization is wrong: we're setting subtree ranges, not individual distances.
	// Let me redo this: first set all to 0, then add edge weights.
	// Reset ft
	ft = NewFenwick(n)

	// Initialize: for each node, add edgeWeight to its subtree
	var init func(u int)
	init = func(u int) {
		if u != 0 {
			ft.rangeAdd(tin[u], tout[u], edgeW[u])
		}
		for _, nb := range adj[u] {
			v := nb[0]
			if v == edgeParent[u] {
				continue
			}
			init(v)
		}
	}
	init(0)

	dist := func(u int) int64 {
		return ft.sum(tin[u])
	}

  // Alokasi slice
	result := make([]int64, 0, len(queries))
	for _, q := range queries {
		if q[0] == 0 {
			// Query distance between u and v
			u, v := q[1], q[2]
			l := lca(u, v)
			d := dist(u) + dist(v) - 2*dist(l)
			result = append(result, d)
		} else {
			// Update edge weight: u, v, new_weight
			u, v, w := q[1], q[2], q[3]
			// Determine which is the child (deeper node)
			child := u
			if depth[v] > depth[u] {
				child = v
			}
			delta := int64(w) - edgeW[child]
			edgeW[child] = int64(w)
			// Update subtree of child
			ft.rangeAdd(tin[child], tout[child], delta)
		}
	}

	return result
}

func main() {
	// Test: n=3, edges=[[0,1,1],[1,2,2]], queries=[[0,0,2]] -> expected [3]
	res := shortestPathWeightedTree(3, [][]int{{0, 1, 1}, {1, 2, 2}}, [][]int{{0, 0, 2}})
	fmt.Printf("Tree(3, edges=[0-1:1, 1-2:2]), query(0,2) -> %v (expected [3])\n", res)

	// Test with update
	res2 := shortestPathWeightedTree(5, [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 4}, {0, 4, 5}},
		[][]int{{0, 0, 2}, {1, 1, 2, 5}, {0, 0, 2}})
	// Initial: 0-2 dist = 0-1-2 = 2+3=5
	// Update edge 1-2 to 5: 0-2 dist = 2+5=7
	fmt.Printf("Tree+update -> %v\n", res2)
}
```
