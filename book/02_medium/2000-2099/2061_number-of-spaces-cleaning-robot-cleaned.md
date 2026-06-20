# 2061 — Number Of Spaces Cleaning Robot Cleaned

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numberOfCleanRooms(room [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2061: Number of Spaces Cleaning Robot Cleaned
// https://leetcode.com/problems/number-of-spaces-cleaning-robot-cleaned/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func numberOfCleanRooms(room [][]int) int {
	m, n := len(room), len(room[0])
  // Matriks 2D
	visited := make([][][4]bool, m)
  // Range loop
	for i := range visited {
		visited[i] = make([][4]bool, n)
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up
  // HashMap: O(1) lookup
	cleaned := make(map[[2]int]bool)
	dir := 0
	r, c := 0, 0
	cleaned[[2]int{0, 0}] = true

	for {
		if visited[r][c][dir] {
			break
		}
		visited[r][c][dir] = true

		// Try to move in current direction
		nextR, nextC := r+dirs[dir][0], c+dirs[dir][1]

		if nextR >= 0 && nextR < m && nextC >= 0 && nextC < n && room[nextR][nextC] == 0 {
			r, c = nextR, nextC
			cleaned[[2]int{r, c}] = true
		} else {
			dir = (dir + 1) % 4
		}
	}

	return len(cleaned)
}

func main() {
	// Test case 1
	room1 := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}
	fmt.Println("Test 1:", numberOfCleanRooms(room1))
	// Expected: 7

	// Test case 2
	room2 := [][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 2:", numberOfCleanRooms(room2))
	// Expected: 1

	// Test case 3
	room3 := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 3:", numberOfCleanRooms(room3))
	// Expected: 9
}
```
