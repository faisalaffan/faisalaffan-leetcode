# 1772 — Sort Features By Popularity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortFeatures(features []string, responses []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(f * r), Space: O(f)  
**Kompleksitas Ruang:** O(f)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
	featureRank := make(map[string]int)
	for i, f := range features {
		featureRank[f] = i
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	for _, resp := range responses {
  // Membuat map (HashMap) — pencarian O(1)
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
