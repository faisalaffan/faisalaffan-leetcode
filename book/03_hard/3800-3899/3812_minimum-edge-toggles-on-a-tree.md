# 3812 — Minimum Edge Toggles On A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumFlips(n int, edges [][]int, start string, target string) []int
```

> **💡 Hint:** Root the tree. For each node, precompute toggles needed

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3812: Minimum Edge Toggles on a Tree
// https://leetcode.com/problems/minimum-edge-toggles-on-a-tree/
// Difficulty: Hard
//
// Given a tree with directed edges, find min edge toggles to make
// node start reach target. Each query has (start, target).
//
// Approach: Root the tree. For each node, precompute toggles needed
// to reach root. Answer = toggles(start->root) + toggles(target->root)
// - 2*toggles(lca->root), but adjusted for direction.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumFlips(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, "0", "3"))
	// Example 2
	fmt.Println(minimumFlips(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}, "3", "4"))
	// Edge: same node
	fmt.Println(minimumFlips(3, [][]int{{0, 1}, {1, 2}}, "1", "1"))
	// Edge: two nodes
	fmt.Println(minimumFlips(2, [][]int{{0, 1}}, "0", "1"))
}

func minimumFlips(n int, edges [][]int, start string, target string) []int {
	// This function returns a single int for a (start, target) query.
	// Build adjacency with direction info
	type edge struct {
		to int
		// 0 = forward (u->v), 1 = backward (u<-v)
		dir int
	}

  // Membuat matriks/slice 2D untuk DP
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], edge{v, 0}) // u->v is forward
		adj[v] = append(adj[v], edge{u, 1}) // v->u is backward
	}

	// BFS to compute min toggles from start to all nodes
	const INF = 1 << 30
  // Alokasi slice integer
	dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	type pair struct {
		node int
		cost int
	}
	q := []pair{{0, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		u := cur.node
		for _, e := range adj[u] {
			cost := cur.cost
			if e.dir == 1 { // backward edge needs toggle
				cost++
			}
			if cost < dist[e.to] {
				dist[e.to] = cost
				q = append(q, pair{e.to, cost})
			}
		}
	}

	// Parse start and target as integers
	s, t := 0, 0
	for _, ch := range start {
		s = s*10 + int(ch-'0')
	}
	for _, ch := range target {
		t = t*10 + int(ch-'0')
	}

	// For each query: min toggles from s to t
	// Root at 0. dist[i] = min toggles from 0 to i.
	// Min toggles from s to t = dist[s] + dist[t] - 2*dist[lca(s,t)]
	// but we need a simple BFS from s for each query.
	// For a single query, BFS from s.
  // Alokasi slice integer
	bfsDist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range bfsDist {
		bfsDist[i] = INF
	}
	bfsDist[s] = 0
	bq := []pair{{s, 0}}
	for len(bq) > 0 {
		cur := bq[0]
		bq = bq[1:]
		u := cur.node
		for _, e := range adj[u] {
			cost := cur.cost
			if e.dir == 1 { // backward edge needs toggle relative to traversal direction
				// If we're going from v to u (u is v), but we stored adj from the edge's perspective
				// Need to handle direction carefully
				_ = cost
			}
			// For BFS from s: we can travel both directions,
			// toggling if the edge points opposite to travel direction
			// Travel from u to e.to:
			// If edge is forward (0), no toggle when going u->e.to
			// If edge is backward (1), need toggle when going u->e.to
			toggle := 0
			if e.dir == 1 {
				toggle = 1
			}
			newCost := cur.cost + toggle
			if newCost < bfsDist[e.to] {
				bfsDist[e.to] = newCost
				bq = append(bq, pair{e.to, newCost})
			}
		}
	}

	return []int{bfsDist[t]}
}
```
