# 0126 — Word Ladder Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findLadders(beginWord string, endWord string, wordList []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DFS, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #126: Word Ladder II
// https://leetcode.com/problems/word-ladder-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return [][]string{}
	}

	// BFS to find shortest distances from beginWord
	dist := map[string]int{beginWord: 0}
	queue := []string{beginWord}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currDist := dist[curr]
		if curr == endWord {
			break
		}
		neighbors := getNeighbors(curr, wordSet)
		for _, nb := range neighbors {
			if _, seen := dist[nb]; !seen {
				dist[nb] = currDist + 1
				queue = append(queue, nb)
			}
		}
	}

	if _, reached := dist[endWord]; !reached {
		return [][]string{}
	}

	// DFS to reconstruct all shortest paths
	result := [][]string{}
	path := []string{beginWord}

	var dfs func(string)
	dfs = func(word string) {
		if word == endWord {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for _, nb := range getNeighbors(word, wordSet) {
			if d, ok := dist[nb]; ok && d == dist[word]+1 {
				path = append(path, nb)
				dfs(nb)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(beginWord)
	return result
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

	result := findLadders(beginWord, endWord, wordList)
	fmt.Printf("findLadders(%q, %q, %v) = %v\n", beginWord, endWord, wordList, result)

	// Expected: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
	expectedCount := 2
	if len(result) == expectedCount {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d paths, got %d\n", expectedCount, len(result))
	}
}
```
