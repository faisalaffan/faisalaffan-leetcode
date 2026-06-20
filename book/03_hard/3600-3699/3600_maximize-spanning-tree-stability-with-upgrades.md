# 3600 — Maximize Spanning Tree Stability With Upgrades

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxStability(n int, edges [][]int, k int) int
```

> **💡 Hint:** Binary search on answer. For a given x, check if we can build a

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3600: Maximize Spanning Tree Stability with Upgrades
// https://leetcode.com/problems/maximize-spanning-tree-stability-with-upgrades/
// Difficulty: Hard
//
// Given n nodes and edges [u,v,strength,must] where must=1 means edge cannot
// be upgraded, maximize the minimum edge strength in a spanning tree after
// upgrading at most k edges (double their strength).
//
// Approach: Binary search on answer. For a given x, check if we can build a
// spanning tree where every edge has strength >= x after at most k upgrades.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(maxStability(4, [][]int{{0, 1, 3, 0}, {1, 2, 5, 1}, {2, 3, 2, 0}, {0, 3, 4, 0}}, 1))
	// Example 2
	fmt.Println(maxStability(3, [][]int{{0, 1, 2, 0}, {1, 2, 1, 0}}, 0))
	// Edge: single node
	fmt.Println(maxStability(1, [][]int{}, 0))
	// Edge: impossible
	fmt.Println(maxStability(4, [][]int{{0, 1, 1, 1}}, 5))
}

func maxStability(n int, edges [][]int, k int) int {
	if n <= 1 {
		return 0
	}

	// Extract unique strengths for binary search range
  // Alokasi slice integer
	strengths := make([]int, 0)
	for _, e := range edges {
		strengths = append(strengths, e[2])
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(strengths)

	// Binary search
	left, right := 0, len(strengths)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if canBuild(n, edges, k, strengths[mid]) {
			result = strengths[mid]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func canBuild(n int, edges [][]int, k int, minStrength int) bool {
	// DSU
  // Alokasi slice integer
	parent := make([]int, n)
  // Alokasi slice integer
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	usedUpgrades := 0
	connected := 0

	// First, add all mandatory edges (must=1) that already meet minStrength
	// Then add optional edges that already meet minStrength
	// Then try upgrading optional edges to meet minStrength

	for _, e := range edges {
		u, v, s, must := e[0], e[1], e[2], e[3]
		if must == 1 && s >= minStrength {
			if find(u) != find(v) {
				union(u, v)
				connected++
			}
		}
	}

	for _, e := range edges {
		u, v, s, must := e[0], e[1], e[2], e[3]
		if must == 0 && s >= minStrength {
			if find(u) != find(v) {
				union(u, v)
				connected++
			}
		}
	}

	for _, e := range edges {
		if usedUpgrades >= k {
			break
		}
		u, v, s, must := e[0], e[1], e[2], e[3]
		if must == 0 && s*2 >= minStrength && s < minStrength {
			if find(u) != find(v) {
				union(u, v)
				connected++
				usedUpgrades++
			}
		}
	}

	return connected == n-1
}
```
