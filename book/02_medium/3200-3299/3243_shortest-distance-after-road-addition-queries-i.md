# 3243 — Shortest Distance After Road Addition Queries I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func shortestDistanceAfterQueries(n int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(q * (n + q))  |  **Ruang:** O(n + q)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3243: Shortest Distance After Road Addition Queries I
// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/
// Difficulty: Medium
// Time: O(q * (n + q)) | Space: O(n + q)

import (
	"fmt"
	"math"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
  // Matriks 2D
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

  // Alokasi slice
	ans := make([]int, len(queries))

	for qi, q := range queries {
		u, v := q[0], q[1]
		graph[u] = append(graph[u], v)

  // Alokasi slice
		dist := make([]int, n)
  // Range loop
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0
		queue := []int{0}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, nb := range graph[cur] {
				if dist[cur]+1 < dist[nb] {
					dist[nb] = dist[cur] + 1
					queue = append(queue, nb)
				}
			}
		}
		ans[qi] = dist[n-1]
	}
	return ans
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}})) // Expected: [3, 2, 1]
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))          // Expected: [1, 1]
}
```
