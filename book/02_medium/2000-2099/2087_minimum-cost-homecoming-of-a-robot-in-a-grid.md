# 2087 — Minimum Cost Homecoming Of A Robot In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m + n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2087: Minimum Cost Homecoming of a Robot in a Grid
// https://leetcode.com/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)

import "fmt"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	cost := 0
	r1, c1 := startPos[0], startPos[1]
	r2, c2 := homePos[0], homePos[1]

	// Move rows
	step := 1
	if r1 > r2 {
		step = -1
	}
	for r := r1 + step; r != r2+step; r += step {
		cost += rowCosts[r]
	}

	// Move columns
	step = 1
	if c1 > c2 {
		step = -1
	}
	for c := c1 + step; c != c2+step; c += step {
		cost += colCosts[c]
	}

	return cost
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCost([]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}))
	// Expected: 18

	// Test case 2
	fmt.Println("Test 2:", minCost([]int{0, 0}, []int{0, 0}, []int{5}, []int{5}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minCost([]int{2, 2}, []int{0, 0}, []int{1, 2, 3}, []int{4, 5, 6}))
	// Expected: 15 (row 1 + row 0 + col 1 + col 0 = 2+1+5+4 = 12... let me recalculate)
	// Moving from row 2 to 0: rowCosts[1] + rowCosts[0] = 2 + 1 = 3
	// Moving from col 2 to 0: colCosts[1] + colCosts[0] = 5 + 4 = 9
	// Total = 12
}
```
