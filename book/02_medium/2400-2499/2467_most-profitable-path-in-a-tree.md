# 2467 — Most Profitable Path In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func mostProfitablePath(edges [][]int, bob int, amount []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2467: Most Profitable Path in a Tree
// https://leetcode.com/problems/most-profitable-path-in-a-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Bob moves to root (fixed path). Alice moves from root to leaf,
// collecting max profit considering time-shared nodes.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	// 6

	fmt.Println(mostProfitablePath([][]int{{0, 1}}, 1, []int{-7280, 2350}))
	// -7280
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Bob's path from bob to 0
  // Alokasi slice integer
	parent := make([]int, n)
	var dfsParent func(u, p int)
	dfsParent = func(u, p int) {
		parent[u] = p
		for _, v := range graph[u] {
			if v != p {
				dfsParent(v, u)
			}
		}
	}
	dfsParent(0, -1)

	// Bob's arrival time at each node
  // Alokasi slice integer
	bobTime := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range bobTime {
		bobTime[i] = math.MaxInt32
	}
	t := 0
	for u := bob; u != -1; u = parent[u] {
		bobTime[u] = t
		t++
	}

	ans := math.MinInt32
	var dfsAlice func(u, p, t, profit int)
	dfsAlice = func(u, p, t, profit int) {
		if t < bobTime[u] {
			profit += amount[u]
		} else if t == bobTime[u] {
			profit += amount[u] / 2
		}
		isLeaf := true
		for _, v := range graph[u] {
			if v != p {
				isLeaf = false
				dfsAlice(v, u, t+1, profit)
			}
		}
		if isLeaf && profit > ans {
			ans = profit
		}
	}
	dfsAlice(0, -1, 0, 0)
	return ans
}
```
