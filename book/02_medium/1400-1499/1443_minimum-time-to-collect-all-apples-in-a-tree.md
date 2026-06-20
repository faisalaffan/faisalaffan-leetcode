# 1443 — Minimum Time To Collect All Apples In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minTime(n int, edges [][]int, hasApple []bool) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(n) for adjacency list and recursion stack

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1443: Minimum Time to Collect All Apples in a Tree
// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, true, true, false})) // 8

	// Test case 2
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, false, true, false})) // 6

	// Test case 3
	fmt.Println(minTime(4, [][]int{{0, 2}, {0, 3}, {1, 2}},
		[]bool{false, true, false, false})) // 4
}

// Time: O(n) where n = number of nodes
// Space: O(n) for adjacency list and recursion stack
func minTime(n int, edges [][]int, hasApple []bool) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(node int) int {
		visited[node] = true
		time := 0

		for _, child := range adj[node] {
			if !visited[child] {
				childTime := dfs(child)
				if childTime > 0 || hasApple[child] {
					time += childTime + 2 // 2 for going down and back up
				}
			}
		}

		return time
	}

	return dfs(0)
}
```
