# 2282 — Number Of People That Can Be Seen In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func visiblePeople(heights [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2282: Number of People That Can Be Seen in a Grid
// https://leetcode.com/problems/number-of-people-that-can-be-seen-in-a-grid/
// Difficulty: Medium [Paid]
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func visiblePeople(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
  // Matriks 2D
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	// For each cell, count visible people to the right
	for i := 0; i < m; i++ {
		stack := []int{}
		for j := n - 1; j >= 0; j-- {
			visible := 0
			// Pop shorter people
			for len(stack) > 0 && stack[len(stack)-1] < heights[i][j] {
				stack = stack[:len(stack)-1]
				visible++
			}
			if len(stack) > 0 {
				visible++
			}
			result[i][j] = visible
			// Push current height
			for len(stack) > 0 && stack[len(stack)-1] == heights[i][j] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, heights[i][j])
		}
	}

	// For each cell, count visible people below
	for j := 0; j < n; j++ {
		stack := []int{}
		for i := m - 1; i >= 0; i-- {
			visible := 0
			for len(stack) > 0 && stack[len(stack)-1] < heights[i][j] {
				stack = stack[:len(stack)-1]
				visible++
			}
			if len(stack) > 0 {
				visible++
			}
			result[i][j] += visible
			for len(stack) > 0 && stack[len(stack)-1] == heights[i][j] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, heights[i][j])
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(visiblePeople([][]int{{3, 1, 4, 2, 5}}))
	// Expected: [[1, 1, 1, 1, 0]]

	// Test case 2
	fmt.Println(visiblePeople([][]int{{5, 1}, {3, 1}, {4, 1}}))
	// Expected: [[1, 0], [1, 0], [0, 0]]
}
```
