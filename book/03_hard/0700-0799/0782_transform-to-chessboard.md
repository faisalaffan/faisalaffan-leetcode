# 0782 — Transform To Chessboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func movesToChessboard(board [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #782: Transform to Chessboard
// https://leetcode.com/problems/transform-to-chessboard/
// Difficulty: Hard
//
// For an n x n board, determine minimum number of row/column swaps to transform
// it into a valid chessboard. A valid chessboard has alternating 0/1 in both
// rows and columns. Validity requires:
//   - Each row equals either row 0 or its complement (otherwise column swaps
//     can never fix it).
//   - Column 0 must have roughly equal 0s and 1s (within 1 for odd n).
//   - Same for row 0.
// Swap count is computed by examining how many first-column entries are already
// in their correct positions for an alternating pattern.

import "fmt"

func main() {
	// Example from problem: returns 2
	fmt.Println(movesToChessboard([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
	}))

	// Already a chessboard (2x2 alternating): should be 0
	fmt.Println(movesToChessboard([][]int{
		{0, 1},
		{1, 0},
	}))

	// Invalid board (impossible): should be -1
	fmt.Println(movesToChessboard([][]int{
		{0, 1},
		{0, 1},
	}))

	// Single cell: 0
	fmt.Println(movesToChessboard([][]int{
		{0},
	}))

	// 3x3 chessboard: should be 0
	fmt.Println(movesToChessboard([][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}))
}

func movesToChessboard(board [][]int) int {
	n := len(board)

	// Check validity: every cell must satisfy the relation
	// board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j] == 0
	// This ensures rows/cols are either identical or complementary.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}

	// Count sum of first column and first row for balance check
	rowSum, colSum := 0, 0
	for i := 0; i < n; i++ {
		rowSum += board[i][0]
		colSum += board[0][i]
	}

	// Balance: each color must appear either n/2 or (n+1)/2 times
	if n%2 == 0 {
		if rowSum*2 != n || colSum*2 != n {
			return -1
		}
	} else {
		if abs(rowSum*2-n) > 1 || abs(colSum*2-n) > 1 {
			return -1
		}
	}

	// Count positions in column 0 / row 0 that already match the 0,1,0,1,...
	// pattern starting with 0 at index 0.
	rowSwap, colSwap := 0, 0
	for i := 0; i < n; i++ {
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}
	}

	// For odd n, only one pattern is possible.
	if n%2 == 1 {
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
	} else {
		rowSwap = min(rowSwap, n-rowSwap)
		colSwap = min(colSwap, n-colSwap)
	}

	// Each swap fixes two misplaced rows (or columns), so divide by 2.
	return (rowSwap + colSwap) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
