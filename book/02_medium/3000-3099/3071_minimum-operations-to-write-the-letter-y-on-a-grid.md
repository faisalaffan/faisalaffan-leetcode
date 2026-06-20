# 3071 — Minimum Operations To Write The Letter Y On A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToWriteY(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3071: Minimum Operations to Write the Letter Y on a Grid
// https://leetcode.com/problems/minimum-operations-to-write-the-letter-y-on-a-grid/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumOperationsToWriteY([][]int{{1, 2, 2}, {2, 2, 3}, {2, 3, 3}}))
	fmt.Println(minimumOperationsToWriteY([][]int{{0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}}))
}

func minimumOperationsToWriteY(grid [][]int) int {
	n := len(grid)
	// Find max value
	maxV := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxV {
				maxV = grid[i][j]
			}
		}
	}
	size := maxV + 1
  // Alokasi slice integer
	yCnt := make([]int, size)
  // Alokasi slice integer
	notYCnt := make([]int, size)
	center := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			isY := false
			if i == j && i <= center {
				isY = true
			} else if i+j == n-1 && i <= center {
				isY = true
			} else if j == center && i >= center {
				isY = true
			}
			if isY {
				yCnt[v]++
			} else {
				notYCnt[v]++
			}
		}
	}
	totalY := 0
	totalNotY := 0
	for _, c := range yCnt {
		totalY += c
	}
	for _, c := range notYCnt {
		totalNotY += c
	}
	ans := n * n
	for yv := 0; yv < size; yv++ {
		for nv := 0; nv < size; nv++ {
			if yv == nv {
				continue
			}
			ops := totalY - yCnt[yv] + totalNotY - notYCnt[nv]
			if ops < ans {
				ans = ops
			}
		}
	}
	return ans
}
```
