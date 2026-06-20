# 0127 — Word Ladder

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ladderLength(beginWord string, endWord string, wordList []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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

  // Membuat map (HashMap) — pencarian O(1)
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
  // Loop linear O(n): iterasi setiap elemen
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
