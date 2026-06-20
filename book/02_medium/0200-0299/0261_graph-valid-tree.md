# 0261 — Graph Valid Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func validTree(n int, edges [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(V+E), Space: O(V+E)  
**Kompleksitas Ruang:** O(V+E)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #261: Graph Valid Tree
// https://leetcode.com/problems/graph-valid-tree/
// Difficulty: Medium [Paid]
// Time: O(V+E), Space: O(V+E)

import "fmt"

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

  // Alokasi slice integer
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		parent[px] = py
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
	fmt.Println(validTree(4, [][]int{{0, 1}, {2, 3}}))
}
```
