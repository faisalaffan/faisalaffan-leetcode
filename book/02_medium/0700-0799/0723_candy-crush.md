# 0723 — Candy Crush

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func candyCrush(board [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(R * C * max(R, C))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #723: Candy Crush
// https://leetcode.com/problems/candy-crush/
// Difficulty: Medium [Paid]
// Time: O(R * C * max(R, C))
// Space: O(1)

import "fmt"

func main() {
	board := [][]int{
		{110, 5, 112, 113, 114},
		{210, 211, 5, 213, 214},
		{310, 311, 3, 313, 314},
		{410, 411, 412, 5, 414},
		{5, 1, 512, 3, 3},
		{610, 4, 1, 613, 614},
		{710, 1, 2, 713, 714},
		{810, 1, 2, 1, 1},
		{1, 1, 2, 2, 2},
		{4, 1, 4, 4, 1014},
	}
	result := candyCrush(board)
	for _, row := range result {
		fmt.Println(row)
	}
}

func candyCrush(board [][]int) [][]int {
	rows, cols := len(board), len(board[0])

	for {
		crushed := false

		// Mark horizontal crushes
		for r := 0; r < rows; r++ {
			for c := 0; c < cols-2; c++ {
				val := abs(board[r][c])
				if val != 0 && abs(board[r][c+1]) == val && abs(board[r][c+2]) == val {
					board[r][c] = -val
					board[r][c+1] = -val
					board[r][c+2] = -val
					crushed = true
				}
			}
		}

		// Mark vertical crushes
		for r := 0; r < rows-2; r++ {
			for c := 0; c < cols; c++ {
				val := abs(board[r][c])
				if val != 0 && abs(board[r+1][c]) == val && abs(board[r+2][c]) == val {
					board[r][c] = -val
					board[r+1][c] = -val
					board[r+2][c] = -val
					crushed = true
				}
			}
		}

		if !crushed {
			break
		}

		// Gravity: drop candies
		for c := 0; c < cols; c++ {
			writeRow := rows - 1
			for r := rows - 1; r >= 0; r-- {
				if board[r][c] > 0 {
					board[writeRow][c] = board[r][c]
					writeRow--
				}
			}
			for r := writeRow; r >= 0; r-- {
				board[r][c] = 0
			}
		}
	}

	return board
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
