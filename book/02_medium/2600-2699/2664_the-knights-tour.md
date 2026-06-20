# 2664 — The Knights Tour

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func tourOfKnight(m int, n int, r int, c int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(8^(m*n)) worst case  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #2664: The Knight's Tour
// https://leetcode.com/problems/the-knights-tour/
// Difficulty: Medium [Paid]
// Time: O(8^(m*n)) worst case | Space: O(m*n)

import "fmt"

var dirs = [8][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

func tourOfKnight(m int, n int, r int, c int) [][]int {
  // Matriks 2D
	board := make([][]int, m)
  // Range loop
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	var solve func(row, col, step int) bool
	solve = func(row, col, step int) bool {
		if step == m*n {
			return true
		}
		for _, d := range dirs {
			nr, nc := row+d[0], col+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == -1 {
				board[nr][nc] = step
				if solve(nr, nc, step+1) {
					return true
				}
				board[nr][nc] = -1
			}
		}
		return false
	}

	board[r][c] = 0
	solve(r, c, 1)
	return board
}

func main() {
	// Test case 1: 1x1 board
	fmt.Println("Test 1:", tourOfKnight(1, 1, 0, 0))
	// Expected: [[0]]

	// Test case 2
	res := tourOfKnight(3, 4, 0, 0)
	fmt.Println("Test 2 dims:", len(res), "x", len(res[0]))
	// Verify all cells filled
	count := 0
	for _, row := range res {
		for _, v := range row {
			if v >= 0 {
				count++
			}
		}
	}
	fmt.Println("  filled:", count)

	// Test case 3
	res2 := tourOfKnight(5, 5, 2, 2)
	fmt.Println("Test 3 dims:", len(res2), "x", len(res2[0]))
}
```
