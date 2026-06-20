# 1252 — Cells With Odd Values In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func oddCells(m, n int, indices [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m + k)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1252: Cells with Odd Values in a Matrix
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
// Difficulty: Easy
// Time: O(n + m + k) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}})) // 6
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}})) // 0
}

// LeetCode submission: oddCells
func oddCells(m, n int, indices [][]int) int {
  // Alokasi slice integer
	rows := make([]int, m)
  // Alokasi slice integer
	cols := make([]int, n)
	for _, idx := range indices {
		rows[idx[0]]++
		cols[idx[1]]++
	}
	oddRows, oddCols := 0, 0
	for _, v := range rows {
		if v%2 == 1 {
			oddRows++
		}
	}
	for _, v := range cols {
		if v%2 == 1 {
			oddCols++
		}
	}
	return oddRows*(n-oddCols) + (m-oddRows)*oddCols
}
```
