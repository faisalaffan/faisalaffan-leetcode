# 3142 — Check If Grid Satisfies Conditions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfGridSatisfiesConditions(grid [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3142: Check if Grid Satisfies Conditions
// https://leetcode.com/problems/check-if-grid-satisfies-conditions/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: satisfiesConditions
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 0, 2}, {1, 0, 2}})) // true
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 1, 1}, {0, 0, 0}})) // false
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: satisfiesConditions
func CheckIfGridSatisfiesConditions(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Below must be equal
			if i+1 < m && grid[i][j] != grid[i+1][j] {
				return false
			}
			// Right must be different
			if j+1 < n && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
```
