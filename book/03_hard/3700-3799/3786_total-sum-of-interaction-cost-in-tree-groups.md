# 3786 — Total Sum Of Interaction Cost In Tree Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func totalInteraction(n int, edges [][]int, group []int) int64
```

> **💡 Hint:** For each edge, count same-group pairs where one node

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3786: Total Sum of Interaction Cost in Tree Groups
// https://leetcode.com/problems/total-sum-of-interaction-cost-in-tree-groups/
// Difficulty: Hard
//
// For each pair of nodes with the same group label, compute the
// number of edges on their path. Sum all such path lengths.
//
// Approach: For each edge, count same-group pairs where one node
// is on each side. Use DFS + subtree group counts.

import "fmt"

func main() {
	// Example 1
	fmt.Println(totalInteraction(3, [][]int{{0, 1}, {1, 2}}, []int{1, 1, 1}))
	// Example 2
	fmt.Println(totalInteraction(3, [][]int{{0, 1}, {1, 2}}, []int{3, 2, 3}))
	// Edge: no same-group pairs
	fmt.Println(totalInteraction(2, [][]int{{0, 1}}, []int{9, 8}))
	// Edge: single node
	fmt.Println(totalInteraction(1, [][]int{}, []int{5}))
}

func totalInteraction(n int, edges [][]int, group []int) int64 {
	// Build adjacency
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Total count of each group across entire tree
  // Membuat map (HashMap) — pencarian O(1)
	totalGroup := make(map[int]int)
	for _, g := range group {
		totalGroup[g]++
	}

	var ans int64

	// DFS: for each edge, count subtree nodes in each group
	var dfs func(u, parent int) map[int]int
	dfs = func(u, parent int) map[int]int {
  // Membuat map (HashMap) — pencarian O(1)
		cnt := make(map[int]int)
		cnt[group[u]] = 1

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			subCnt := dfs(v, u)

			// For each group in subtree, edge (u,v) contributes
			// subCnt[g] * (totalGroup[g] - subCnt[g])
			for g, c := range subCnt {
				pairCnt := int64(c) * int64(totalGroup[g]-c)
				ans += pairCnt
			}

			// Merge subtree counts
			for g, c := range subCnt {
				cnt[g] += c
			}
		}
		return cnt
	}

	dfs(0, -1)
	return ans
}
```
