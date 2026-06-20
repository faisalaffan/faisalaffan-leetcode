# 1772 — Sort Features By Popularity

## Deskripsi

**Soal:** [1772. Sort Features By Popularity](https://leetcode.com/problems/sort-features-by-popularity/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(f * r), Space: O(f)  
**Kompleksitas Ruang:** O(f)

**Algoritma:** —

**Fungsi Solusi:** `func sortFeatures(features []string, responses []string) []string`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	featureRank := make(map[string]int)
	for i, f := range features {
		featureRank[f] = i
	}

  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	for _, resp := range responses {
  // Membuat map untuk pencarian O(1): key → value
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

  // Membuat slice untuk menyimpan hasil
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
