# 1258 — Synonymous Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func newUF() *uf`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Backtracking, Sorting, Union-Find

**Waktu:** O(2^k) where k = number of synonym groups per sentence  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1258: Synonymous Sentences
// https://leetcode.com/problems/synonymous-sentences/
// Difficulty: Medium [Paid]

// Union-Find to group synonyms, then generate all sentences
// by replacing words with all synonyms in their group.

// Time: O(2^k) where k = number of synonym groups per sentence
// Space: O(n)

type uf struct {
	parent map[string]string
}

func newUF() *uf {
	return &uf{parent: make(map[string]string)}
}

func (u *uf) find(x string) string {
	if _, exists := u.parent[x]; !exists {
		u.parent[x] = x
	}
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) union(x, y string) {
	u.parent[u.find(x)] = u.find(y)
}

func generateSentences(synonyms [][]string, text string) []string {
	u := newUF()
	for _, s := range synonyms {
		u.union(s[0], s[1])
	}

	// Group words by root
  // HashMap: O(1) lookup
	groups := make(map[string][]string)
	for _, s := range synonyms {
		for _, w := range s {
			root := u.find(w)
			groups[root] = append(groups[root], w)
		}
	}

	// Sort and dedupe each group
	for root := range groups {
  // HashMap: O(1) lookup
		wordSet := make(map[string]bool)
		for _, w := range groups[root] {
			wordSet[w] = true
		}
		groups[root] = make([]string, 0, len(wordSet))
		for w := range wordSet {
			groups[root] = append(groups[root], w)
		}
		sort.Strings(groups[root])
	}

	// Map word -> group root
  // HashMap: O(1) lookup
	wordToRoot := make(map[string]string)
	for root, words := range groups {
		for _, w := range words {
			wordToRoot[w] = root
		}
	}

	words := strings.Fields(text)
	result := make([]string, 0)

	var backtrack func(idx int, current []string)
	backtrack = func(idx int, current []string) {
		if idx == len(words) {
			result = append(result, strings.Join(current, " "))
			return
		}
		root, hasSynonyms := wordToRoot[words[idx]]
		if !hasSynonyms {
			backtrack(idx+1, append(current, words[idx]))
		} else {
			for _, syn := range groups[root] {
				backtrack(idx+1, append(current, syn))
			}
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	result := generateSentences([][]string{{"happy", "joy"}, {"sad", "sorrow"}}, "I am happy today")
	fmt.Printf("%v\n", result)
}
```
