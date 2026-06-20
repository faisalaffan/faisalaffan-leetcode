# 0425 — Word Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordSquares(words []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking, Trie, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #425: Word Squares
// https://leetcode.com/problems/word-squares/
// Difficulty: Hard
//
// Trie + backtracking. Build a trie of all words, then backtrack to form squares
// row by row. At each step, the next word must have a prefix matching the
// characters already placed in the current column.

func main() {
	// Example 1: ["area","lead","wall","lady","ball"] => 2 squares
	result := wordSquares([]string{"area", "lead", "wall", "lady", "ball"})
	fmt.Println("Word squares:", result)
	// Example 2: ["abat","baba","atan","atal"] => 2 squares
	result = wordSquares([]string{"abat", "baba", "atan", "atal"})
	fmt.Println("Word squares:", result)
	// Single word
	result = wordSquares([]string{"abc"})
	fmt.Println("Single:", result)
	// Edge: no solution
	result = wordSquares([]string{"a", "b"})
	fmt.Println("No solution:", result)
}

type trieNode struct {
	children [26]*trieNode
	words    []string
}

func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}
	n := len(words[0])
  // Edge case: input kosong
	if n == 0 {
		return nil
	}

	root := &trieNode{}
	for _, w := range words {
		node := root
		root.words = append(root.words, w)
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
			node.words = append(node.words, w)
		}
	}

	var result [][]string
	var current []string

	var backtrack func(step int)
	backtrack = func(step int) {
		if step == n {
			square := make([]string, n)
			copy(square, current)
			result = append(result, square)
			return
		}

		prefix := make([]byte, step)
		for i := 0; i < step; i++ {
			prefix[i] = current[i][step]
		}

		node := root
		for _, ch := range prefix {
			idx := ch - 'a'
			if node.children[idx] == nil {
				return
			}
			node = node.children[idx]
		}

		for _, candidate := range node.words {
			current = append(current, candidate)
			backtrack(step + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
```
