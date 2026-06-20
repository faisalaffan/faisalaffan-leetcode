# 0699 — Falling Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func fallingSquares(positions [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #699: Falling Squares
// https://leetcode.com/problems/falling-squares/
// Difficulty: Hard
//
// For each square, check overlap with all previous squares.
// Base height = max height of overlapping squares below.
// Current height = base + size. Track running max.

func main() {
	// [[1,2],[2,3],[6,1]] => [2,5,5]
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	// [[100,100],[200,100]] => [100,100]
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}}))
	// [[2,1],[2,9],[1,8]] => [1,10,18]
	fmt.Println(fallingSquares([][]int{{2, 1}, {2, 9}, {1, 8}}))
	// Single
	fmt.Println(fallingSquares([][]int{{5, 5}}))
	// [[1,2],[3,4]] => [2,4]
	fmt.Println(fallingSquares([][]int{{1, 2}, {3, 4}}))
}

func fallingSquares(positions [][]int) []int {
	n := len(positions)
  // Alokasi slice integer
	ans := make([]int, n)
  // Alokasi slice integer
	heights := make([]int, n)

	for i, p := range positions {
		left, size := p[0], p[1]
		right := left + size

		base := 0
		for j := 0; j < i; j++ {
			jLeft := positions[j][0]
			jRight := jLeft + positions[j][1]
			if left < jRight && right > jLeft {
				if heights[j] > base {
					base = heights[j]
				}
			}
		}
		heights[i] = base + size

		if i == 0 {
			ans[i] = heights[i]
		} else {
			if heights[i] > ans[i-1] {
				ans[i] = heights[i]
			} else {
				ans[i] = ans[i-1]
			}
		}
	}

	return ans
}
```
