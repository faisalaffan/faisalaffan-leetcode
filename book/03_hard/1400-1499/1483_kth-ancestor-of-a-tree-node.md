# 1483 — Kth Ancestor Of A Tree Node

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(n int, parent []int) TreeAncestor`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Bitmask** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1483: Kth Ancestor of a Tree Node
// https://leetcode.com/problems/kth-ancestor-of-a-tree-node/
// Difficulty: Hard
//
// Approach: Binary Lifting (doubling)
// up[node][i] = 2^i-th ancestor of node.
// up[node][0] = parent[node]
// up[node][i] = up[up[node][i-1]][i-1]
// getKthAncestor: for each bit of k, jump up.

import "fmt"

func main() {
	// Example: n=7, parent=[-1,0,0,1,1,2,2]
	ta := Constructor(7, []int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(ta.GetKthAncestor(3, 1)) // 1
	fmt.Println(ta.GetKthAncestor(5, 2)) // 0
	fmt.Println(ta.GetKthAncestor(6, 3)) // -1

	// Edge case
	ta2 := Constructor(1, []int{-1})
	fmt.Println(ta2.GetKthAncestor(0, 1)) // -1
}

type TreeAncestor struct {
	up [][]int // up[node][i]
}

func Constructor(n int, parent []int) TreeAncestor {
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}
  // Matriks 2D
	up := make([][]int, n)
	for i := 0; i < n; i++ {
		up[i] = make([]int, LOG)
		up[i][0] = parent[i]
	}
	for j := 1; j < LOG; j++ {
		for i := 0; i < n; i++ {
			if up[i][j-1] >= 0 {
				up[i][j] = up[up[i][j-1]][j-1]
			} else {
				up[i][j] = -1
			}
		}
	}
	return TreeAncestor{up: up}
}

func (ta *TreeAncestor) GetKthAncestor(node int, k int) int {
	LOG := len(ta.up[0])
	for j := 0; j < LOG; j++ {
		if k&(1<<j) != 0 {
			node = ta.up[node][j]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}
```
