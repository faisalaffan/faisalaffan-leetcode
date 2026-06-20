# 0064 — Minimum Path Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func minPathSum(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #64: Minimum Path Sum
// https://leetcode.com/problems/minimum-path-sum/
// Difficulty: Medium

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7

	// Test case 2
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})) // 12

	// Test case 3
	fmt.Println(minPathSum([][]int{{1}})) // 1
}

// Time: O(m*n) | Space: O(1)
```
