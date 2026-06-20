# 1245 — Tree Diameter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func treeDiameter(edges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1245: Tree Diameter
// https://leetcode.com/problems/tree-diameter/
// Difficulty: Medium [Paid]

// Two BFS: find farthest node from any node, then farthest from that node.
// Distance = diameter.

// Time: O(n)
// Space: O(n)

func treeDiameter(edges [][]int) int {
	if len(edges) == 0 {
		return 0
	}

	n := len(edges) + 1
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	bfs := func(start int) (int, int) {
  // Alokasi slice integer
		dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dist {
			dist[i] = -1
		}
		queue := []int{start}
		dist[start] = 0
		farthestNode, maxDist := start, 0

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
					if dist[v] > maxDist {
						maxDist = dist[v]
						farthestNode = v
					}
				}
			}
		}
		return farthestNode, maxDist
	}

	farthest, _ := bfs(0)
	_, diameter := bfs(farthest)
	return diameter
}

func main() {
	fmt.Printf("%d (expected: 2)\n", treeDiameter([][]int{{0, 1}, {1, 2}, {2, 3}}))
	fmt.Printf("%d (expected: 3)\n", treeDiameter([][]int{{0, 1}, {0, 2}, {1, 3}}))
	fmt.Printf("%d (expected: 0)\n", treeDiameter([][]int{}))
}
```
