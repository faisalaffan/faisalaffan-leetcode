# 2946 — Matrix Similarity After Cyclic Shifts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func MatrixSimilarityAfterCyclicShifts(mat [][]int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2946: Matrix Similarity After Cyclic Shifts
// https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areSimilar
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 4)) // false
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2)) // true
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{2, 2}, {2, 2}}, 3)) // true
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: areSimilar
func MatrixSimilarityAfterCyclicShifts(mat [][]int, k int) bool {
	for _, row := range mat {
		m := len(row)
		if m == 0 {
			continue
		}
		shift := k % m
		if shift == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			// After left shift by k, row[(j + k) % m] moves to position j
			if row[j] != row[(j+shift)%m] {
				return false
			}
		}
	}
	return true
}
```
