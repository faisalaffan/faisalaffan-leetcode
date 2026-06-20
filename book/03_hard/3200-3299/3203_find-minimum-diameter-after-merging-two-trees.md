# 3203 — Find Minimum Diameter After Merging Two Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDiameterAfterMergingTwoTrees(edges1, edges2 [][]int) int
```

> **💡 Hint:** compute diameter via double-BFS (or DFS) for each tree.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3203: Find Minimum Diameter After Merging Two Trees
// https://leetcode.com/problems/find-minimum-diameter-after-merging-two-trees/
// Difficulty: Hard
//
// Given two trees (undirected acyclic graphs), connect one node from each with
// an edge. Find the minimum possible diameter of the resulting tree.
//
// Answer = max(d1, d2, ceil(d1/2)+ceil(d2/2)+1).
//
// Approach: compute diameter via double-BFS (or DFS) for each tree.

import "fmt"

func minimumDiameterAfterMergingTwoTrees(edges1, edges2 [][]int) int {
	d1 := treeDiameter(edges1)
	d2 := treeDiameter(edges2)
	merge := (d1+1)/2 + (d2+1)/2 + 1

	ans := d1
	if d2 > ans {
		ans = d2
	}
	if merge > ans {
		ans = merge
	}
	return ans
}

func treeDiameter(edges [][]int) int {
	n := len(edges) + 1
	if n <= 1 {
		return 0
	}
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS from 0 to find farthest node.
	far1, _ := bfs(adj, 0)
	// BFS from farthest node to get diameter.
	_, dist := bfs(adj, far1)
	return dist
}

func bfs(adj [][]int, start int) (farthest, maxDist int) {
	n := len(adj)
  // Alokasi slice integer
	dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = -1
	}
	q := []int{start}
	dist[start] = 0
	farthest = start
	maxDist = 0

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				q = append(q, v)
				if dist[v] > maxDist {
					maxDist = dist[v]
					farthest = v
				}
			}
		}
	}
	return farthest, maxDist
}

func main() {
	edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	edges2 := [][]int{{0, 1}}
	fmt.Println(minimumDiameterAfterMergingTwoTrees(edges1, edges2))
}
```
