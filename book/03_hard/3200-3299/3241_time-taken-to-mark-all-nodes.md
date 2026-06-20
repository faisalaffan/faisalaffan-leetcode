# 3241 — Time Taken To Mark All Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func timeTaken(edges [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3241: Time Taken to Mark All Nodes
// https://leetcode.com/problems/time-taken-to-mark-all-nodes/
// Difficulty: Hard
//
// Rerooting DP. First DFS computes the longest marking time into the subtree
// for each node. Second DFS reroots to compute the answer for every node.
//
// Marking rule (from problem description):
// - Odd-indexed node gets marked 1 time unit after an adjacent node is marked.
// - Even-indexed node gets marked 2 time units after an adjacent node is marked.
// Equivalently: propagating from u to v takes 1 if v is odd, 2 if v is even.

import "fmt"

func main() {
	// Example 1: n=4, edges=[[0,1],[0,2],[1,3]] => [2,4,4,5]
	fmt.Println(timeTaken([][]int{{0, 1}, {0, 2}, {1, 3}}))
	// Example 2: n=3, edges=[[0,1],[0,2]] => [2,4,3]
	fmt.Println(timeTaken([][]int{{0, 1}, {0, 2}}))
	// Example 3: n=2, edges=[[0,1]] => [1,2]
	fmt.Println(timeTaken([][]int{{0, 1}}))
	// Example 4: from problem description
	fmt.Println(timeTaken([][]int{{2, 4}, {0, 1}, {2, 3}, {0, 2}}))
	// Example 5: single node
	fmt.Println(timeTaken([][]int{}))
}

func timeTaken(edges [][]int) []int {
	n := len(edges) + 1
  // Edge case: input kosong — langsung return
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []int{0}
	}

  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// dp1[u] = max cost from u into its subtree
	// dp2[u] = second max cost from u into its subtree (for rerooting)
  // Alokasi slice integer
	dp1 := make([]int, n)
  // Alokasi slice integer
	dp2 := make([]int, n)

	// Cost to propagate from parent to child v (depends on target node v)
	propagateCost := func(v int) int {
		if v&1 == 1 { // v is odd
			return 1
		}
		return 2
	}

	var dfs1 func(u, p int) int
	dfs1 = func(u, p int) int {
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			cost := propagateCost(v) + dfs1(v, u)
			if cost > dp1[u] {
				dp2[u] = dp1[u]
				dp1[u] = cost
			} else if cost > dp2[u] {
				dp2[u] = cost
			}
		}
		return dp1[u]
	}
	dfs1(0, -1)

  // Alokasi slice integer
	ans := make([]int, n)

	var dfs2 func(u, p, other int)
	dfs2 = func(u, p, other int) {
		ans[u] = dp1[u]
		if other > ans[u] {
			ans[u] = other
		}

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			// Cost to propagate from v up to u (target is u)
			upCost := propagateCost(u)
			var newOther int
			if dp1[v]+propagateCost(v) == dp1[u] {
				newOther = other
				if dp2[u] > newOther {
					newOther = dp2[u]
				}
			} else {
				newOther = other
				if dp1[u] > newOther {
					newOther = dp1[u]
				}
			}
			newOther += upCost
			dfs2(v, u, newOther)
		}
	}
	dfs2(0, -1, 0)

	return ans
}
```
