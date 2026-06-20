# 1258 — Synonymous Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func newUF() *uf
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking, Union-Find (DSU)

**Kompleksitas Waktu:** O(2^k) where k = number of synonym groups per sentence  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[string][]string)
	for _, s := range synonyms {
		for _, w := range s {
			root := u.find(w)
			groups[root] = append(groups[root], w)
		}
	}

	// Sort and dedupe each group
	for root := range groups {
  // Membuat map (HashMap) — pencarian O(1)
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
  // Membuat map (HashMap) — pencarian O(1)
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
