# 1181 — Before And After Puzzle

## Deskripsi

**Soal:** [1181. Before And After Puzzle](https://leetcode.com/problems/before-and-after-puzzle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2 * L)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func beforeAndAfterPuzzles(phrases []string) []string`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1181: Before and After Puzzle
// https://leetcode.com/problems/before-and-after-puzzle/
// Difficulty: Medium [Paid]

// Given list of phrases. Merge phrase i and j if last word of i
// equals first word of j. Result: i + j[firstWordLen:].
// Return sorted unique results.

// Time: O(n^2 * L)
// Space: O(n^2)

func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)
  // Membuat slice untuk menyimpan hasil
	firstWords := make([]string, n)
  // Membuat slice untuk menyimpan hasil
	lastWords := make([]string, n)
  // Membuat slice 2D untuk DP/tabel
	words := make([][]string, n)

	for i, p := range phrases {
		words[i] = strings.Fields(p)
		firstWords[i] = words[i][0]
		lastWords[i] = words[i][len(words[i])-1]
	}

  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[string]bool)
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if lastWords[i] == firstWords[j] {
				merged := phrases[i] + phrases[j][len(firstWords[j]):]
				if !seen[merged] {
					seen[merged] = true
					result = append(result, merged)
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [writing code rocks])\n",
		beforeAndAfterPuzzles([]string{"writing code", "code rocks"}))

	fmt.Printf("%v (expected: [a d a b c d])\n",
		beforeAndAfterPuzzles([]string{"a b", "b c", "c d"}))
}
```
