# 0425 — Word Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordSquares(words []string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking, Trie, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Edge case: input kosong — langsung return
	if n == 0 {
		return nil
	}

	root := &trieNode{}
	for _, w := range words {
		node := root
		root.words = append(root.words, w)
  // Loop linear O(n): iterasi setiap elemen
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
