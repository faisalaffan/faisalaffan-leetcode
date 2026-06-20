# 3402 — Minimum Operations To Make Columns Strictly Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeColumnsStrictlyIncreasing(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3402: Minimum Operations to Make Columns Strictly Increasing
// https://leetcode.com/problems/minimum-operations-to-make-columns-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}}))
}

// MinimumOperationsToMakeColumnsStrictlyIncreasing returns minimum operations to make each column strictly increasing.
// Each operation increments an element by 1.
// Time: O(n * m). Space: O(1).
func MinimumOperationsToMakeColumnsStrictlyIncreasing(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	ops := 0
	for c := 0; c < cols; c++ {
		for r := 1; r < rows; r++ {
			if grid[r][c] <= grid[r-1][c] {
				diff := grid[r-1][c] - grid[r][c] + 1
				ops += diff
				grid[r][c] += diff
			}
		}
	}
	return ops
}
```
