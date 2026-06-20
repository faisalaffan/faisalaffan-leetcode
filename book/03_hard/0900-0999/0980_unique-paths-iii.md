# 0980 — Unique Paths Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func uniquePathsIII(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #980: Unique Paths III
// https://leetcode.com/problems/unique-paths-iii/
// Difficulty: Hard

import "fmt"

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 {
		return 0
	}

	startX, startY := 0, 0
	nonObstacleCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
			}
			if grid[i][j] != -1 {
				nonObstacleCount++
			}
		}
	}

  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	ans := 0
	var dfs func(x, y, walked int)
	dfs = func(x, y, walked int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == -1 || visited[x][y] {
			return
		}
		if grid[x][y] == 2 {
			if walked == nonObstacleCount {
				ans++
			}
			return
		}

		visited[x][y] = true
		dfs(x-1, y, walked+1)
		dfs(x+1, y, walked+1)
		dfs(x, y-1, walked+1)
		dfs(x, y+1, walked+1)
		visited[x][y] = false
	}

	dfs(startX, startY, 1)
	return ans
}

func main() {
	fmt.Println("Example 1:")
	grid1 := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
	fmt.Println(uniquePathsIII(grid1))
	// Expected: 2

	fmt.Println("Example 2:")
	grid2 := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}
	fmt.Println(uniquePathsIII(grid2))
	// Expected: 4

	fmt.Println("Example 3:")
	grid3 := [][]int{
		{0, 1},
		{2, 0},
	}
	fmt.Println(uniquePathsIII(grid3))
	// Expected: 0
}
```
