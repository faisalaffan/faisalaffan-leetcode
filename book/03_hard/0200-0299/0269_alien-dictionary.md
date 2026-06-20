# 0269 — Alien Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func alienOrder(words []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #269: Alien Dictionary
// https://leetcode.com/problems/alien-dictionary/
// Difficulty: Hard [Paid]

import (
	"fmt"
)

func alienOrder(words []string) string {
  // HashMap: O(1) lookup
	graph := make(map[byte][]byte)
  // HashMap: O(1) lookup
	inDegree := make(map[byte]int)

	// Initialize all characters
	for _, w := range words {
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			ch := w[i]
			if _, exists := graph[ch]; !exists {
				graph[ch] = make([]byte, 0)
				inDegree[ch] = 0
			}
		}
	}

	// Build graph
  // Linear scan O(n)
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}

		// Check for invalid prefix case: w1 is longer and w1 starts with w2
		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}

		for j := 0; j < minLen; j++ {
			c1, c2 := w1[j], w2[j]
			if c1 != c2 {
				graph[c1] = append(graph[c1], c2)
				inDegree[c2]++
				break
			}
		}
	}

	// Topological sort (BFS / Kahn's algorithm)
	queue := make([]byte, 0)
	for ch := range graph {
		if inDegree[ch] == 0 {
			queue = append(queue, ch)
		}
	}

	result := make([]byte, 0, len(graph))
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		result = append(result, ch)

		for _, neighbor := range graph[ch] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != len(graph) {
		return ""
	}

	return string(result)
}

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}
```
