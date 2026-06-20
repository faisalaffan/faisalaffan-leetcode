# 0127 — Word Ladder

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ladderLength(beginWord string, endWord string, wordList []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #127: Word Ladder
// https://leetcode.com/problems/word-ladder/
// Difficulty: Hard

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	// Bidirectional BFS
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := map[string]bool{beginWord: true, endWord: true}

	dist := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		// Always expand the smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

  // HashMap: O(1) lookup
		nextSet := make(map[string]bool)
		for word := range beginSet {
			neighbors := getNeighbors(word, wordSet)
			for _, nb := range neighbors {
				if endSet[nb] {
					return dist + 1
				}
				if !visited[nb] {
					visited[nb] = true
					nextSet[nb] = true
				}
			}
		}

		beginSet = nextSet
		dist++
	}

	return 0
}

func getNeighbors(word string, wordSet map[string]bool) []string {
	neighbors := []string{}
	bytes := []byte(word)
  // Linear scan O(n)
	for i := 0; i < len(bytes); i++ {
		original := bytes[i]
		for c := 'a'; c <= 'z'; c++ {
			bytes[i] = byte(c)
			candidate := string(bytes)
			if candidate != word && wordSet[candidate] {
				neighbors = append(neighbors, candidate)
			}
		}
		bytes[i] = original
	}
	return neighbors
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	result := ladderLength(beginWord, endWord, wordList)
	expected := 5

	fmt.Printf("ladderLength(%q, %q, %v) = %d\n", beginWord, endWord, wordList, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
