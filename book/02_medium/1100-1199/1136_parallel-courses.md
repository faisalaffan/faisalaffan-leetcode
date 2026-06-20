# 1136 — Parallel Courses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumSemesters(n int, relations [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n + len(relations))  |  **Ruang:** O(n + len(relations))

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1136: Parallel Courses
// https://leetcode.com/problems/parallel-courses/
// Difficulty: Medium [Paid]

// topological sort (Kahn's algorithm) to find minimum semesters.

// Time: O(n + len(relations))
// Space: O(n + len(relations))

func minimumSemesters(n int, relations [][]int) int {
  // Matriks 2D
	adj := make([][]int, n+1)
  // Alokasi slice
	indeg := make([]int, n+1)

	for _, r := range relations {
		adj[r[0]] = append(adj[r[0]], r[1])
		indeg[r[1]]++
	}

  // Alokasi slice
	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	semesters := 0
	taken := 0

	for len(queue) > 0 {
		semesters++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			taken++
			for _, next := range adj[cur] {
				indeg[next]--
				if indeg[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	if taken != n {
		return -1
	}
	return semesters
}

func main() {
	fmt.Printf("%d (expected: 2)\n", minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Printf("%d (expected: -1)\n", minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	fmt.Printf("%d (expected: 1)\n", minimumSemesters(3, [][]int{}))
}
```
