# 0212 — Word Search Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func findWords(board [][]byte, words []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Trie

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #212: Word Search II
// https://leetcode.com/problems/word-search-ii/
// Difficulty: Hard

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}

	result := make([]string, 0)
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		ch := board[r][c]
		if ch == '#' {
			return
		}
		idx := ch - 'a'
		child := node.children[idx]
		if child == nil {
			return
		}
		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		board[r][c] = '#'
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) {
				dfs(nr, nc, child)
			}
		}
		board[r][c] = ch
	}

  // Range loop: iterasi dengan indeks + nilai
	for i := range board {
		for j := range board[0] {
			dfs(i, j, root)
		}
	}
	return result
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}
```
