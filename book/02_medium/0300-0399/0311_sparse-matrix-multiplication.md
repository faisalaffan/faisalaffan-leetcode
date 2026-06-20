# 0311 — Sparse Matrix Multiplication

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func multiply(mat1 [][]int, mat2 [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n*k), Space: O(m*k) optimized for sparse matrices  
**Kompleksitas Ruang:** O(m*k) optimized for sparse matrices

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #311: Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/
// Difficulty: Medium [Paid]
// Time: O(m*n*k), Space: O(m*k) optimized for sparse matrices

import "fmt"

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m, k, n := len(mat1), len(mat1[0]), len(mat2[0])
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for kk := 0; kk < k; kk++ {
			if mat1[i][kk] != 0 {
				for j := 0; j < n; j++ {
					if mat2[kk][j] != 0 {
						result[i][j] += mat1[i][kk] * mat2[kk][j]
					}
				}
			}
		}
	}

	return result
}

func main() {
	mat1 := [][]int{{1, 0, 0}, {-1, 0, 3}}
	mat2 := [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{0}}
	mat2 = [][]int{{0}}
	fmt.Println(multiply(mat1, mat2))

	mat1 = [][]int{{1, -5}}
	mat2 = [][]int{{12}, {-1}}
	fmt.Println(multiply(mat1, mat2))
}
```
