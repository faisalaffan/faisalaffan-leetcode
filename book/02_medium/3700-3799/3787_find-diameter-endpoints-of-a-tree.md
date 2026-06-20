# 3787 — Find Diameter Endpoints Of A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func findDiameterEndpointsOfATree(n int, edges [][]int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3787: Find Diameter Endpoints of a Tree
// https://leetcode.com/problems/find-diameter-endpoints-of-a-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func findDiameterEndpointsOfATree(n int, edges [][]int) string {
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	bfs := func(start int) (int, []int) {
  // Alokasi slice integer
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[start] = 0
		q := []int{start}
		far := start
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			if dist[u] > dist[far] {
				far = u
			}
			for _, v := range g[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
		return far, dist
	}

	a, _ := bfs(0)
	b, distA := bfs(a)
	_, distB := bfs(b)
	diameter := distA[b]

	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		if distA[i] == diameter || distB[i] == diameter {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(findDiameterEndpointsOfATree(7, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {3, 5}, {1, 6}}))
	fmt.Println(findDiameterEndpointsOfATree(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findDiameterEndpointsOfATree(1, [][]int{}))
}
```
