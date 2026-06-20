# 3565 — Sequential Grid Path Cover

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func SequentialGridPathCover(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3565: Sequential Grid Path Cover
// https://leetcode.com/problems/sequential-grid-path-cover/
// Difficulty: Medium [Paid]
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Test 1:", SequentialGridPathCover(grid))
	// Test case 2
	grid2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Test 2:", SequentialGridPathCover(grid2))
	// Test case 3
	grid3 := [][]int{{1}}
	fmt.Println("Test 3:", SequentialGridPathCover(grid3))
}

func SequentialGridPathCover(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	// Find if there's a path visiting all cells in sequential order
	// (1, 2, 3, ..., m*n)
	prev := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if prev != -1 && grid[i][j] != prev+1 {
				return 0
			}
			prev = grid[i][j]
		}
	}
	return 1
}
```
