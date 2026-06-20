# 2123 — Minimum Operations To Remove Adjacent Ones In Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToRemoveAdjacentOnes(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2123: Minimum Operations to Remove Adjacent Ones in Matrix
// https://leetcode.com/problems/minimum-operations-to-remove-adjacent-ones-in-matrix/
// Difficulty: Hard [Paid]
//
// Minimum vertex cover in a bipartite graph. Each operation at (i,j) toggles
// the cell and its 4 neighbors. Model as: for each adjacent pair of 1s, at
// least one must be flipped. This reduces to max bipartite matching (Kőnig's
// theorem). Graph cells are colored like a chessboard based on (i+j) parity.

import "fmt"

func main() {
	// Example
	grid1 := [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid1))

	// All 1s diagonal - no adjacent 1s
	grid2 := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid2))

	// All adjacent
	grid3 := [][]int{{1, 1}, {1, 1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid3))

	// Single cell
	grid4 := [][]int{{1}}
	fmt.Println(minimumOperationsToRemoveAdjacentOnes(grid4))
}

func minimumOperationsToRemoveAdjacentOnes(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Assign IDs to cells with value 1
  // Membuat matriks/slice 2D untuk DP
	id := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range id {
		id[i] = make([]int, n)
		for j := range id[i] {
			id[i][j] = -1
		}
	}

	leftCount := 0
	rightCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if (i+j)%2 == 0 {
					id[i][j] = leftCount
					leftCount++
				} else {
					id[i][j] = rightCount
					rightCount++
				}
			}
		}
	}

	if leftCount == 0 || rightCount == 0 {
		return 0
	}

	// Build adjacency from left to right
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, leftCount)
	for i := 0; i < leftCount; i++ {
		adj[i] = []int{}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (i+j)%2 == 0 {
				u := id[i][j]
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
						v := id[ni][nj]
						adj[u] = append(adj[u], v)
					}
				}
			}
		}
	}

	// Maximum bipartite matching using DFS (Kuhn's algorithm)
  // Alokasi slice integer
	matchR := make([]int, rightCount)
  // Range loop: iterasi dengan indeks + nilai
	for i := range matchR {
		matchR[i] = -1
	}

	var dfs func(u int, seen []bool) bool
	dfs = func(u int, seen []bool) bool {
		for _, v := range adj[u] {
			if seen[v] {
				continue
			}
			seen[v] = true
			if matchR[v] == -1 || dfs(matchR[v], seen) {
				matchR[v] = u
				return true
			}
		}
		return false
	}

	result := 0
	for u := 0; u < leftCount; u++ {
		seen := make([]bool, rightCount)
		if dfs(u, seen) {
			result++
		}
	}

	return result
}
```
