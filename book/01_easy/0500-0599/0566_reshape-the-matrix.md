# 0566 — Reshape The Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReshapeTheMatrix(mat [][]int, r, c int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #566: Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(m*n)
func ReshapeTheMatrix(mat [][]int, r, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, r)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}

func main() {
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(ReshapeTheMatrix([][]int{{1, 2}, {3, 4}}, 2, 4))
}
```
