# 0059 — Spiral Matrix Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func generateMatrix(n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #59: Spiral Matrix II
// https://leetcode.com/problems/spiral-matrix-ii/
// Difficulty: Medium

import "fmt"

func generateMatrix(n int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	matrix := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		if top <= bottom {
			for j := right; j >= left; j-- {
				matrix[bottom][j] = num
				num++
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}

func main() {
	// Test case 1
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]

	// Test case 2
	fmt.Println(generateMatrix(1)) // [[1]]

	// Test case 3
	fmt.Println(generateMatrix(4))
}

// Time: O(n^2) | Space: O(n^2)
```
