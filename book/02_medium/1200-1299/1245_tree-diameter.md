# 1245 — Tree Diameter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func treeDiameter(edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

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
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	bfs := func(start int) (int, int) {
  // Alokasi slice
		dist := make([]int, n)
  // Range loop
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
