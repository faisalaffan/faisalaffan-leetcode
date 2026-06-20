# 2858 — Minimum Edge Reversals So Every Node Is Reachable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func minEdgeReversals(n int, edges [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2858: Minimum Edge Reversals So Every Node Is Reachable
// https://leetcode.com/problems/minimum-edge-reversals-so-every-node-is-reachable/
// Difficulty: Hard
//
// Directed graph whose underlying undirected structure is a tree of n nodes.
// For each node i, find minimum edge reversals needed so every node is
// reachable from i. Rerooting DP:
//   - First DFS from node 0: count reversals needed, building weights
//     (0 = forward edge, 1 = backward edge/reversal needed).
//   - Second DFS (reroot): propagate answer to children.
//     If edge u->v with cost 0, answer[v] = answer[u] + 1 (now backward).
//     If edge u->v with cost 1, answer[v] = answer[u] - 1 (now forward).
// O(N) time, O(N) space.

import "fmt"

func minEdgeReversals(n int, edges [][]int) []int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], [2]int{v, 0}) // forward: no reversal
		adj[v] = append(adj[v], [2]int{u, 1}) // backward: reversal needed
	}

  // Alokasi slice integer
	ans := make([]int, n)

	// First DFS: compute reversals needed if starting from node 0
	var dfs1 func(u, p int)
	dfs1 = func(u, p int) {
		for _, nei := range adj[u] {
			v, cost := nei[0], nei[1]
			if v == p {
				continue
			}
			ans[0] += cost
			dfs1(v, u)
		}
	}
	dfs1(0, -1)

	// Second DFS: reroot the answer
	var dfs2 func(u, p int)
	dfs2 = func(u, p int) {
		for _, nei := range adj[u] {
			v, cost := nei[0], nei[1]
			if v == p {
				continue
			}
			if cost == 0 {
				// Edge was u->v (forward). When rerooting at v, this edge
				// becomes v->u which needs reversal: +1
				ans[v] = ans[u] + 1
			} else {
				// Edge was v->u (had cost 1 from v). When rerooting at v,
				// this edge is now correctly oriented: -1
				ans[v] = ans[u] - 1
			}
			dfs2(v, u)
		}
	}
	dfs2(0, -1)

	return ans
}

func main() {
	// Example 1: n=4, edges=[[2,0],[2,1],[1,3]] => [1,1,0,2]
	res1 := minEdgeReversals(4, [][]int{{2, 0}, {2, 1}, {1, 3}})
	for _, v := range res1 {
		fmt.Println(v)
	}

	// Example 2: n=3, edges=[[1,2],[2,0]] => [2,0,1]
	res2 := minEdgeReversals(3, [][]int{{1, 2}, {2, 0}})
	for _, v := range res2 {
		fmt.Println(v)
	}

	// Single edge
	fmt.Println(minEdgeReversals(2, [][]int{{0, 1}}))

	// Reverse direction
	fmt.Println(minEdgeReversals(2, [][]int{{1, 0}}))

	// Star topology: center 0 connected to all
	fmt.Println(minEdgeReversals(4, [][]int{{0, 1}, {0, 2}, {0, 3}}))

	// All edges away from 0
	fmt.Println(minEdgeReversals(4, [][]int{{0, 1}, {1, 2}, {2, 3}}))

	// All edges toward 0
	fmt.Println(minEdgeReversals(4, [][]int{{1, 0}, {2, 1}, {3, 2}}))
}
```
