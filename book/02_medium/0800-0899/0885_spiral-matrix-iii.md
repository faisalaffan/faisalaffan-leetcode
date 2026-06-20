# 0885 — Spiral Matrix Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func SpiralMatrixIii(rows int, cols int, rStart int, cStart int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(rows * cols)  
**Kompleksitas Ruang:** O(rows * cols)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #885: Spiral Matrix III
// https://leetcode.com/problems/spiral-matrix-iii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SpiralMatrixIii(1, 4, 0, 0))
	fmt.Println(SpiralMatrixIii(5, 6, 1, 4))
}

// Time: O(rows * cols) | Space: O(rows * cols)
func SpiralMatrixIii(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
  // Membuat matriks/slice 2D untuk DP
	ans := make([][]int, 0, total)
	ans = append(ans, []int{rStart, cStart})
	if total == 1 {
		return ans
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := rStart, cStart
	step := 1
	dir := 0

	for len(ans) < total {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				r += dirs[dir][0]
				c += dirs[dir][1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					ans = append(ans, []int{r, c})
				}
			}
			dir = (dir + 1) % 4
		}
		step++
	}

	return ans
}
```
