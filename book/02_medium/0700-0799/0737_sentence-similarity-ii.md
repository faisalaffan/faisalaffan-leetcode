# 0737 — Sentence Similarity Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Union-Find

**Waktu:** O(n * alpha(n))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #737: Sentence Similarity II
// https://leetcode.com/problems/sentence-similarity-ii/
// Difficulty: Medium [Paid]
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	pairs := [][]string{
		{"great", "fine"},
		{"drama", "acting"},
		{"fine", "good"},
	}
	fmt.Println(areSentencesSimilarTwo([]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, pairs))
}

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}

  // HashMap: O(1) lookup
	parent := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

  // Linear scan O(n)
	for i := 0; i < len(words1); i++ {
		if words1[i] == words2[i] {
			continue
		}
		if find(words1[i]) != find(words2[i]) {
			return false
		}
	}

	return true
}
```
