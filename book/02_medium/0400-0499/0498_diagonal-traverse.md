# 0498 — Diagonal Traverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DiagonalTraverse(mat [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #498: Diagonal Traverse
// https://leetcode.com/problems/diagonal-traverse/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(DiagonalTraverse([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(DiagonalTraverse([][]int{{1, 2}, {3, 4}}))
}

func DiagonalTraverse(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])
  // Alokasi slice integer
	result := make([]int, m*n)
	row, col := 0, 0
	dir := 1 // 1 = up-right, -1 = down-left

	for i := 0; i < m*n; i++ {
		result[i] = mat[row][col]
		if dir == 1 { // moving up-right
			if col == n-1 {
				row++
				dir = -1
			} else if row == 0 {
				col++
				dir = -1
			} else {
				row--
				col++
			}
		} else { // moving down-left
			if row == m-1 {
				col++
				dir = 1
			} else if col == 0 {
				row++
				dir = 1
			} else {
				row++
				col--
			}
		}
	}

	return result
}
```
