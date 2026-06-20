# 1886 — Determine Whether Matrix Can Be Obtained By Rotation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindRotation(mat [][]int, target [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1886: Determine Whether Matrix Can Be Obtained by Rotation
// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
// Difficulty: Easy

import "fmt"

// Time: O(n^2), Space: O(1)
func FindRotation(mat [][]int, target [][]int) bool {
	for rotation := 0; rotation < 4; rotation++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
	}
	return false
}

func rotate(mat [][]int) [][]int {
	n := len(mat)
  // Membuat matriks/slice 2D untuk DP
	rotated := make([][]int, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = mat[n-1-j][i]
		}
	}
	return rotated
}

func equal(a, b [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}))
}
```
