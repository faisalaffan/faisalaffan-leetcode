# 0766 — Toeplitz Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func isToeplitzMatrix(matrix [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(m*n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #766: Toeplitz Matrix
// https://leetcode.com/problems/toeplitz-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})) // true
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))                            // false
}

// isToeplitzMatrix checks if every diagonal from top-left to bottom-right has the same element.
// Time: O(m*n). Space: O(1).
func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
```
