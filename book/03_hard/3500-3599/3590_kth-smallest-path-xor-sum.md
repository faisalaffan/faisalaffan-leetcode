# 3590 — Kth Smallest Path Xor Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func kthSmallest(parent []int, vals []int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Binary Search, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3590: Kth Smallest Path XOR Sum
// https://leetcode.com/problems/kth-smallest-path-xor-sum/
// Difficulty: Hard
//
// Given a tree with parent array and values, for each query (u, k), find the
// k-th smallest distinct path XOR sum in the subtree of node u.
//
// Approach: DFS to compute root-to-node XOR. For each subtree, collect distinct
// XOR values using small-to-large merging. For queries, binary search on sorted
// XOR values.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(kthSmallest([]int{-1, 0, 0}, []int{1, 2, 3}, [][]int{{0, 1}, {1, 1}, {2, 1}}))
	// Example 2
	fmt.Println(kthSmallest([]int{-1, 0, 1, 1}, []int{1, 2, 3, 4}, [][]int{{1, 2}}))
	// Edge: single node
	fmt.Println(kthSmallest([]int{-1}, []int{5}, [][]int{{0, 1}}))
	// Edge: k out of range
	fmt.Println(kthSmallest([]int{-1, 0}, []int{1, 2}, [][]int{{0, 10}}))
}

func kthSmallest(parent []int, vals []int, queries [][]int) []int {
	n := len(parent)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

  // Matriks 2D
	children := make([][]int, n)
	root := -1
	for i := 0; i < n; i++ {
		if parent[i] == -1 {
			root = i
		} else {
			children[parent[i]] = append(children[parent[i]], i)
		}
	}

	// Compute root-to-node XOR
  // Alokasi slice
	xorToRoot := make([]int, n)
	var dfsXor func(u int, x int)
	dfsXor = func(u int, x int) {
		x ^= vals[u]
		xorToRoot[u] = x
		for _, v := range children[u] {
			dfsXor(v, x)
		}
	}
	dfsXor(root, 0)

	// For each node, collect distinct XOR values in its subtree
	// Map node -> sorted list of distinct XOR values
  // Alokasi slice
	subtreeXors := make([]map[int]bool, n)
	var dfsCollect func(u int) map[int]bool
	dfsCollect = func(u int) map[int]bool {
  // HashMap: O(1) lookup
		set := make(map[int]bool)
		set[xorToRoot[u]] = true
		for _, v := range children[u] {
			childSet := dfsCollect(v)
			// Small-to-large merging
			if len(childSet) > len(set) {
				set, childSet = childSet, set
			}
			for x := range childSet {
				set[x] = true
			}
		}
		subtreeXors[u] = set
		return set
	}
	dfsCollect(root)

	// Sort subtree XOR values for binary search
  // Matriks 2D
	sortedXors := make([][]int, n)
	for i := 0; i < n; i++ {
		for x := range subtreeXors[i] {
			sortedXors[i] = append(sortedXors[i], x)
		}
  // Sort O(n log n)
		sort.Ints(sortedXors[i])
	}

  // Alokasi slice
	result := make([]int, len(queries))
	for qi, q := range queries {
		u, k := q[0], q[1]
		if k <= 0 || k > len(sortedXors[u]) {
			result[qi] = -1
		} else {
			result[qi] = sortedXors[u][k-1]
		}
	}

	return result
}
```
