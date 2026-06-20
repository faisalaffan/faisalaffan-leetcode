# 0192 — Word Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordFrequency(text string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #192: Word Frequency
// https://leetcode.com/problems/word-frequency/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func wordFrequency(text string) []string {
	words := strings.Fields(text)
  // HashMap: O(1) lookup
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++
	}

	type kv struct {
		word  string
		count int
	}

	var sorted []kv
	for w, c := range freq {
		sorted = append(sorted, kv{w, c})
	}

  // Custom sort
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count != sorted[j].count {
			return sorted[i].count > sorted[j].count
		}
		return sorted[i].word < sorted[j].word
	})

	result := make([]string, len(sorted))
	for i, kv := range sorted {
		result[i] = fmt.Sprintf("%s %d", kv.word, kv.count)
	}
	return result
}

func main() {
	for _, line := range wordFrequency("the day is sunny the the the sunny is is") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("hello world hello") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("") {
		fmt.Println(line)
	}
}
```
