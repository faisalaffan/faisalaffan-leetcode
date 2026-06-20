# 1293 — Shortest Path In A Grid With Obstacles Elimination

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestPath(grid [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1293: Shortest Path in a Grid with Obstacles Elimination
// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1293. Shortest Path in a Grid with Obstacles Elimination")
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}
	fmt.Println("k=1:", shortestPath(grid, 1), "(expected 6)")

	grid2 := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}
	fmt.Println("k=1:", shortestPath(grid2, 1), "(expected -1)")
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// visited[r][c][e] = true if visited (r,c) with e eliminations used.
  // Membuat matriks/slice 2D untuk DP
	visited := make([][][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}

	type state struct {
		r, c, elim, dist int
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := []state{{0, 0, 0, 0}}
	visited[0][0][0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.r == m-1 && cur.c == n-1 {
			return cur.dist
		}

		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			newElim := cur.elim
			if grid[nr][nc] == 1 {
				newElim++
			}
			if newElim > k {
				continue
			}

			if !visited[nr][nc][newElim] {
				visited[nr][nc][newElim] = true
				queue = append(queue, state{nr, nc, newElim, cur.dist + 1})
			}
		}
	}

	return -1
}
```
