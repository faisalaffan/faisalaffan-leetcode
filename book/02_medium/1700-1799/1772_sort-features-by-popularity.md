# 1772 — Sort Features By Popularity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func sortFeatures(features []string, responses []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(f * r), Space: O(f)  |  **Ruang:** O(f)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1772: Sort Features by Popularity
// https://leetcode.com/problems/sort-features-by-popularity/
// Difficulty: Medium [Paid]
// Time: O(f * r), Space: O(f)

import (
	"fmt"
	"sort"
)

func sortFeatures(features []string, responses []string) []string {
  // HashMap: O(1) lookup
	featureRank := make(map[string]int)
	for i, f := range features {
		featureRank[f] = i
	}

  // HashMap: O(1) lookup
	freq := make(map[string]int)
	for _, resp := range responses {
  // HashMap: O(1) lookup
		seen := make(map[string]bool)
		word := ""
		for _, ch := range resp + " " {
			if ch == ' ' {
				if word != "" && !seen[word] {
					freq[word]++
					seen[word] = true
				}
				word = ""
			} else {
				word += string(ch)
			}
		}
	}

	sorted := make([]string, len(features))
	copy(sorted, features)
	sort.SliceStable(sorted, func(i, j int) bool {
		fi, fj := freq[sorted[i]], freq[sorted[j]]
		if fi != fj {
			return fi > fj
		}
		return featureRank[sorted[i]] < featureRank[sorted[j]]
	})
	return sorted
}

func main() {
	fmt.Println(sortFeatures(
		[]string{"cooler", "lock", "touch"},
		[]string{"i like cooler cooler", "lock touch cool", "locker like touch"},
	)) // Expected: ["touch", "cooler", "lock"] or ["touch", "lock", "cooler"] depending on frequency

	fmt.Println(sortFeatures(
		[]string{"a", "b", "c"},
		[]string{"a b", "b c", "c a"},
	)) // Expected: ["a", "b", "c"] (all appear in 2 responses, stable sort by original order)
}
```
