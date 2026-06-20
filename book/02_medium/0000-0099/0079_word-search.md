# 0079 — Word Search

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func exist(board [][]byte, word string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n*4^L)  |  **Ruang:** O(L)


## 💻 Solusi Go

```go
package main

// LeetCode #79: Word Search
// https://leetcode.com/problems/word-search/
// Difficulty: Medium

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word)-1 {
			return board[i][j] == word[idx]
		}
		if board[i][j] != word[idx] {
			return false
		}

		tmp := board[i][j]
		board[i][j] = '#'

		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n {
				if dfs(ni, nj, idx+1) {
					board[i][j] = tmp
					return true
				}
			}
		}

		board[i][j] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	// Test case 1
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED")) // true

	// Test case 2
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "SEE")) // true

	// Test case 3
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCB")) // false
}

// Time: O(m*n*4^L) | Space: O(L)
```
