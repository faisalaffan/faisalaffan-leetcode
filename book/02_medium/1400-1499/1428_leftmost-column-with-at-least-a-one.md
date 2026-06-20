# 1428 — Leftmost Column With At Least A One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m + n) where m = rows, n = cols  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1428: Leftmost Column with at Least a One
// https://leetcode.com/problems/leftmost-column-with-at-least-a-one/
// Difficulty: Medium

import "fmt"

type BinaryMatrix struct {
	grid [][]int
}

func (bm BinaryMatrix) Get(row, col int) int {
	return bm.grid[row][col]
}

func (bm BinaryMatrix) Dimensions() []int {
	if len(bm.grid) == 0 {
		return []int{0, 0}
	}
	return []int{len(bm.grid), len(bm.grid[0])}
}

func main() {
	// Test case 1
	bm := BinaryMatrix{[][]int{{0, 0}, {1, 1}}}
	fmt.Println(leftMostColumnWithOne(bm)) // 0

	// Test case 2
	bm2 := BinaryMatrix{[][]int{{0, 0}, {0, 1}}}
	fmt.Println(leftMostColumnWithOne(bm2)) // 1

	// Test case 3
	bm3 := BinaryMatrix{[][]int{{0, 0}, {0, 0}}}
	fmt.Println(leftMostColumnWithOne(bm3)) // -1
}

// Time: O(m + n) where m = rows, n = cols
// Space: O(1)
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	m, n := dim[0], dim[1]

	row, col := 0, n-1
	leftmost := -1

	for row < m && col >= 0 {
		if binaryMatrix.Get(row, col) == 1 {
			leftmost = col
			col--
		} else {
			row++
		}
	}

	return leftmost
}
```
