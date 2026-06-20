# 0529 — Minesweeper

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func UpdateBoard(board [][]byte, click []int) [][]byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #529: Minesweeper
// https://leetcode.com/problems/minesweeper/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	result := UpdateBoard(board, []int{3, 0})
	for _, row := range result {
		fmt.Println(string(row))
	}
}

func UpdateBoard(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	}
	reveal(board, r, c)
	return board
}

func reveal(board [][]byte, r, c int) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != 'E' {
		return
	}
	m, n := len(board), len(board[0])
	mines := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'M' {
				mines++
			}
		}
	}
	if mines > 0 {
		board[r][c] = byte('0' + mines)
		return
	}
	board[r][c] = 'B'
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			reveal(board, r+dr, c+dc)
		}
	}
}
```
