# 2282 — Number Of People That Can Be Seen In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func visiblePeople(heights [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat matriks/slice 2D untuk DP
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
