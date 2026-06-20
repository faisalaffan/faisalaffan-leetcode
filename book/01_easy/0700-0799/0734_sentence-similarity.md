# 0734 — Sentence Similarity

## Deskripsi

**Soal:** [0734. Sentence Similarity](https://leetcode.com/problems/sentence-similarity/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + p) where n = len(sentence), p = len(pairs). Space: O(p).  
**Kompleksitas Ruang:** O(p).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #734: Sentence Similarity
// https://leetcode.com/problems/sentence-similarity/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	s1 := []string{"great", "acting", "skills"}
	s2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // true

	s1 = []string{"great"}
	s2 = []string{"great"}
	fmt.Println(areSentencesSimilar(s1, s2, [][]string{})) // true

	s1 = []string{"great"}
	s2 = []string{"doubleplus", "good"}
	fmt.Println(areSentencesSimilar(s1, s2, pairs)) // false
}

// areSentencesSimilar checks if two sentences are similar according to given similarity pairs.
// Time: O(n + p) where n = len(sentence), p = len(pairs). Space: O(p).
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	// Build bidirectional map
  // Membuat map untuk pencarian O(1): key → value
	pairMap := make(map[string]map[string]bool)
	for _, p := range similarPairs {
		a, b := p[0], p[1]
		if pairMap[a] == nil {
			pairMap[a] = make(map[string]bool)
		}
		if pairMap[b] == nil {
			pairMap[b] = make(map[string]bool)
		}
		pairMap[a][b] = true
		pairMap[b][a] = true
	}

  // Iterasi seluruh elemen
	for i := range sentence1 {
		w1, w2 := sentence1[i], sentence2[i]
		if w1 == w2 {
			continue
		}
		if !pairMap[w1][w2] {
			return false
		}
	}
	return true
}
```
