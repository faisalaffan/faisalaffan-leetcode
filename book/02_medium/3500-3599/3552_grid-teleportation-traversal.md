# 3552 — Grid Teleportation Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func GridTeleportationTraversal(grid [][]int, start, end []int, teleports [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3552: Grid Teleportation Traversal
// https://leetcode.com/problems/grid-teleportation-traversal/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import (
	"container/list"
	"fmt"
)

func main() {
	// Test case 1
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	teleports := [][]int{{0, 0, 2, 2}}
	fmt.Println("Test 1:", GridTeleportationTraversal(grid, []int{0, 0}, []int{2, 2}, teleports))
	// Test case 2
	grid2 := [][]int{{0, 0}, {0, 0}}
	teleports2 := [][]int{{0, 0, 1, 1}}
	fmt.Println("Test 2:", GridTeleportationTraversal(grid2, []int{0, 0}, []int{0, 1}, teleports2))
	// Test case 3
	grid3 := [][]int{{0}}
	fmt.Println("Test 3:", GridTeleportationTraversal(grid3, []int{0, 0}, []int{0, 0}, [][]int{}))
}

func GridTeleportationTraversal(grid [][]int, start, end []int, teleports [][]int) int {
	m, n := len(grid), len(grid[0])
	// Build teleport map
  // Membuat map (HashMap) — pencarian O(1)
	tpMap := make(map[[2]int][2]int)
	for _, tp := range teleports {
		tpMap[[2]int{tp[0], tp[1]}] = [2]int{tp[2], tp[3]}
	}

	// BFS
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := list.New()
	queue.PushBack([3]int{start[0], start[1], 0})
	visited[start[0]][start[1]] = true

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).([3]int)
		r, c, dist := front[0], front[1], front[2]
		if r == end[0] && c == end[1] {
			return dist
		}
		// Check teleport
		if dest, ok := tpMap[[2]int{r, c}]; ok {
			if !visited[dest[0]][dest[1]] {
				visited[dest[0]][dest[1]] = true
				queue.PushBack([3]int{dest[0], dest[1], dist + 1})
			}
		}
		// Normal movement
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				queue.PushBack([3]int{nr, nc, dist + 1})
			}
		}
	}
	return -1
}
```
