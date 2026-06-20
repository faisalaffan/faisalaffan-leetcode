# 2120 — Execution Of All Suffix Instructions Staying In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func executeInstructions(n int, startPos []int, s string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(m * n) or O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2120: Execution of All Suffix Instructions Staying in a Grid
// https://leetcode.com/problems/execution-of-all-suffix-instructions-staying-in-a-grid/
// Difficulty: Medium
// Time: O(m * n) or O(n^2) | Space: O(n)

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	m := len(s)
  // Alokasi slice
	result := make([]int, m)

	dirs := map[byte][2]int{
		'L': {0, -1},
		'R': {0, 1},
		'U': {-1, 0},
		'D': {1, 0},
	}

	for i := 0; i < m; i++ {
		r, c := startPos[0], startPos[1]
		count := 0
		for j := i; j < m; j++ {
			dr, dc := dirs[s[j]][0], dirs[s[j]][1]
			r += dr
			c += dc
			if r < 0 || r >= n || c < 0 || c >= n {
				break
			}
			count++
		}
		result[i] = count
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", executeInstructions(3, []int{0, 1}, "RRDDLU"))
	// Expected: [1, 5, 4, 3, 1, 0]

	// Test case 2
	fmt.Println("Test 2:", executeInstructions(2, []int{1, 1}, "LURD"))
	// Expected: [4, 1, 0, 0]

	// Test case 3
	fmt.Println("Test 3:", executeInstructions(1, []int{0, 0}, "LRUD"))
	// Expected: [0, 0, 0, 0]
}
```
