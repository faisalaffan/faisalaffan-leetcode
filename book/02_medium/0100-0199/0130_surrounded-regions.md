# 0130 — Surrounded Regions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func solve(board [][]byte) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #130: Surrounded Regions
// https://leetcode.com/problems/surrounded-regions/
// Difficulty: Medium

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// Mark border 'O' cells
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

	// Flip remaining 'O' to 'X' and '#' back to 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	// Test case 1
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board) // [[X X X X] [X X X X] [X X X X] [X O X X]]

	// Test case 2
	board = [][]byte{{'X'}}
	solve(board)
	fmt.Println(board) // [[X]]
}

// Time: O(m*n) | Space: O(m*n)
```
