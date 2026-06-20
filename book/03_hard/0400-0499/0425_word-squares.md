# 0425 — Word Squares

## Deskripsi

**Soal:** [0425. Word Squares](https://leetcode.com/problems/word-squares/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Backtracking (pelacakan mundur), Trie (pohon awalan)

## Solusi Go

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
  // Loop standar: indeks 0 sampai n-1
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
  // Membuat slice untuk menyimpan hasil
			square := make([]string, n)
			copy(square, current)
			result = append(result, square)
			return
		}

  // Membuat slice untuk menyimpan hasil
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
