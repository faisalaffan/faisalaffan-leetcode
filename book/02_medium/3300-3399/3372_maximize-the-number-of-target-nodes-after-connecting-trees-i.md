# 3372 — Maximize The Number Of Target Nodes After Connecting Trees I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n^2 + m^2) Space: O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3372: Maximize the Number of Target Nodes After Connecting Trees I
// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/
// Difficulty: Medium
// Time: O(n^2 + m^2) Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}}, 3))
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {1, 2}, {1, 3}}, [][]int{{0, 1}, {1, 2}, {1, 3}}, 2))
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	g1 := buildGraph3372(edges1, n)
	g2 := buildGraph3372(edges2, m)

	maxFrom2 := 0
	for i := 0; i < m; i++ {
		cnt := countWithinDist3372(g2, i, k-1, m)
		if cnt > maxFrom2 {
			maxFrom2 = cnt
		}
	}

  // Alokasi slice integer
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = countWithinDist3372(g1, i, k, n) + maxFrom2
	}
	return ans
}

func buildGraph3372(edges [][]int, n int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	return g
}

func countWithinDist3372(g [][]int, start int, maxDist int, n int) int {
	if maxDist < 0 {
		return 0
	}
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true
	dist := 0
	count := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			count++
			for _, v := range g[queue[i]] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
		queue = queue[size:]
		dist++
		if dist > maxDist {
			break
		}
	}
	return count
}
```
