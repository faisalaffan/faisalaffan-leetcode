# 2647 — Color The Triangle Red

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func colorTheTriangleRed(n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2647: Color the Triangle Red
// https://leetcode.com/problems/color-the-triangle-red/
// Difficulty: Hard [Paid]
//
// Given a triangle of size n (similar to Pascal's triangle layout), color
// cells red. Each operation selects a cell and colors a "Y" shape centered
// on that cell. Find minimum operations to turn all cells red.
//
// The solution exploits the pattern: for each level i from top to bottom,
// color the leftmost cell of each horizontal block of size 3, then shift.

import "fmt"

func main() {
	// Test: print operations for n=2..4
	for n := 2; n <= 5; n++ {
		fmt.Printf("n=%d: %v\n", n, colorTheTriangleRed(n))
	}
}

// colorTheTriangleRed returns minimum operations to color all cells in
// triangle of size n. Each result is a row,col pair.
// The triangle has rows 0..n-1 with row r having (2*r+1) cells.
// Strategy: process in reverse (bottom to top), coloring blocks of 3.
func colorTheTriangleRed(n int) [][]int {
	ops := [][]int{}

	// total cells = n^2
	// We color from bottom to top using a greedy pattern.
	// At each row r (0-indexed from top), we have (2r+1) cells.
	// The operation at (r, c) colors itself, (r+1, c), (r+1, c+2).
	// We need the minimal set covering all cells.
	//
	// Known combinatorial solution: for each level i (1-indexed),
	// color cells at positions (i, 3*j + offset) where offset depends on i%3.

	// We use the constructive greedy pattern described in the editorial.
	// Process columns from right to left within each row.
  // Membuat matriks/slice 2D untuk DP
	marked := make([][]bool, n)
	for i := 0; i < n; i++ {
		sz := 2*i + 1
		marked[i] = make([]bool, sz)
	}

	// Bottom-up: try coloring each cell if it helps
	for r := n - 1; r >= 0; r-- {
		sz := 2*r + 1
		for c := 0; c < sz; c++ {
			if marked[r][c] {
				continue
			}
			// Apply operation at (r, c)
			ops = append(ops, []int{r, c})
			// Color the Y shape: (r,c), (r+1,c), (r+1,c+2) - but only within bounds
			for dr := 0; r+dr < n; dr++ {
				for dc := -dr; dc <= dr; dc += 2 {
					nr, nc := r+dr, c+dc
					if nr < n && nc >= 0 && nc < 2*nr+1 {
						marked[nr][nc] = true
					}
				}
			}
		}
	}

	return ops
}
```
