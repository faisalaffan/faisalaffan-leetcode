# 0079 — Word Search

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func exist(board [][]byte, word string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(m*n*4^L)  
**Kompleksitas Ruang:** O(L)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
