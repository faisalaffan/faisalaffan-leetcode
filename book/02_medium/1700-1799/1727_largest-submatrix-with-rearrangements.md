# 1727 — Largest Submatrix With Rearrangements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestSubmatrix(matrix [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1727: Largest Submatrix With Rearrangements
// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
// Difficulty: Medium
// Time: O(m * n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	maxArea := 0

  // Alokasi slice integer
	heights := make([]int, n)

	for r := 0; r < m; r++ {
		// Update heights
		for c := 0; c < n; c++ {
			if matrix[r][c] == 1 {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}

		// Sort heights for this row (to find max rectangle that can be formed
		// by rearranging columns)
  // Alokasi slice integer
		sorted := make([]int, n)
		copy(sorted, heights)
  // Custom sort dengan comparator
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		})

		// For each column, area = height * (col index) because it's sorted
		for c := 0; c < n; c++ {
			area := sorted[c] * (c + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}})) // Expected: 4
	fmt.Println(largestSubmatrix([][]int{{1, 0, 1, 0, 1}})) // Expected: 3
	fmt.Println(largestSubmatrix([][]int{{1, 1, 0}, {1, 0, 1}})) // Expected: 2
}
```
