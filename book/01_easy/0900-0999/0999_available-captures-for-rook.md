# 0999 — Available Captures For Rook

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numRookCaptures(board [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) (fixed 8x8 board). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #999: Available Captures for Rook
// https://leetcode.com/problems/available-captures-for-rook/
// Difficulty: Easy

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(numRookCaptures(board)) // 3

	board2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(numRookCaptures(board2)) // 0
}

// numRookCaptures counts how many pawns the rook can capture.
// Time: O(1) (fixed 8x8 board). Space: O(1).
func numRookCaptures(board [][]byte) int {
	rRow, rCol := -1, -1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rRow, rCol = i, j
				break
			}
		}
		if rRow != -1 {
			break
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	count := 0
	for _, d := range dirs {
		r, c := rRow+d[0], rCol+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			if board[r][c] == 'B' {
				break
			}
			if board[r][c] == 'p' {
				count++
				break
			}
			r += d[0]
			c += d[1]
		}
	}
	return count
}
```
