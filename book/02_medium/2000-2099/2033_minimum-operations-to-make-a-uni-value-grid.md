# 2033 — Minimum Operations To Make A Uni Value Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(grid [][]int, x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n log(m*n))  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2033: Minimum Operations to Make a Uni-Value Grid
// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
// Difficulty: Medium
// Time: O(m*n log(m*n)) | Space: O(m*n)

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice integer
	vals := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals = append(vals, grid[i][j])
		}
	}

	// Check all values have same remainder mod x
	rem := vals[0] % x
	for _, v := range vals {
		if v%x != rem {
			return -1
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(vals)
	median := vals[len(vals)/2]

	ops := 0
	for _, v := range vals {
		diff := v - median
		if diff < 0 {
			diff = -diff
		}
		ops += diff / x
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([][]int{{2, 4}, {6, 8}}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minOperations([][]int{{1, 5}, {2, 3}}, 1))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", minOperations([][]int{{1, 2}, {3, 4}}, 2))
	// Expected: -1
}
```
