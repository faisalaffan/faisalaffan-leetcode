# 0310 — Minimum Height Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findMinHeightTrees(n int, edges [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(V), Space: O(V+E)  
**Kompleksitas Ruang:** O(V+E)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #310: Minimum Height Trees
// https://leetcode.com/problems/minimum-height-trees/
// Difficulty: Medium
// Time: O(V), Space: O(V+E)

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

  // Alokasi slice integer
	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		adj[edge[0]][edge[1]] = true
		adj[edge[1]][edge[0]] = true
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeaves := []int{}

		for _, leaf := range leaves {
			for neighbor := range adj[leaf] {
				delete(adj[neighbor], leaf)
				if len(adj[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
}
```
