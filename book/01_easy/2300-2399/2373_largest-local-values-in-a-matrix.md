# 2373 — Largest Local Values In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestLocalValuesInAMatrix(grid [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2373: Largest Local Values in a Matrix
// https://leetcode.com/problems/largest-local-values-in-a-matrix/
// Difficulty: Easy
// Time O(n^2) | Space O(n^2)

import "fmt"

func main() {
	fmt.Println(LargestLocalValuesInAMatrix([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}})) // [[9,9],[8,6]]
	fmt.Println(LargestLocalValuesInAMatrix([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})) // [[2,2,2],[2,2,2],[2,2,2]]
}

func LargestLocalValuesInAMatrix(grid [][]int) [][]int {
	n := len(grid)
  // Membuat matriks/slice 2D untuk DP
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			maxVal := 0
			for r := i; r < i+3; r++ {
				for c := j; c < j+3; c++ {
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			res[i][j] = maxVal
		}
	}
	return res
}
```
