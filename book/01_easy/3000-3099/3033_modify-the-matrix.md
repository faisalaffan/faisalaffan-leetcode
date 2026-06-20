# 3033 — Modify The Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func ModifyTheMatrix(matrix [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(1) (modifies in place)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3033: Modify the Matrix
// https://leetcode.com/problems/modify-the-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: modifiedMatrix
	fmt.Println(ModifyTheMatrix([][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}))
	// [[1 2 9] [4 8 6] [7 8 9]]
}

// Time: O(n * m) | Space: O(1) (modifies in place)
// LeetCode submission name: modifiedMatrix
func ModifyTheMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	m, n := len(matrix), len(matrix[0])

	// Find max in each column
  // Alokasi slice integer
	colMax := make([]int, n)
	for j := 0; j < n; j++ {
		maxVal := -1
		for i := 0; i < m; i++ {
			if matrix[i][j] > maxVal {
				maxVal = matrix[i][j]
			}
		}
		colMax[j] = maxVal
	}

	// Replace -1 values
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = colMax[j]
			}
		}
	}
	return matrix
}
```
