# 3913 — Sort Vowels By Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isVowel(ch byte) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(N)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3913: Sort Vowels by Frequency
// https://leetcode.com/problems/sort-vowels-by-frequency/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count vowel frequencies, sort vowels by freq desc (tie: alpha),
// rebuild string by placing sorted vowels in vowel positions.

import (
	"fmt"
	"sort"
)

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func SortVowelsByFrequency(s string) string {
	n := len(s)
  // HashMap: O(1) lookup
	freq := make(map[byte]int)
	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			freq[s[i]]++
		}
	}

	vowels := make([]byte, 0, len(freq))
	for v := range freq {
		vowels = append(vowels, v)
	}
  // Custom sort
	sort.Slice(vowels, func(i, j int) bool {
		if freq[vowels[i]] != freq[vowels[j]] {
			return freq[vowels[i]] > freq[vowels[j]]
		}
		return vowels[i] < vowels[j]
	})

	ans := make([]byte, n)
	vi := 0
	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			ans[i] = vowels[vi]
			freq[vowels[vi]]--
			if freq[vowels[vi]] == 0 {
				vi++
			}
		} else {
			ans[i] = s[i]
		}
	}

	return string(ans)
}

func main() {
	// Example
	fmt.Println(SortVowelsByFrequency("hello")) // "holle" or "hollo"? vowels: e,o. Both freq 1. alpha: e < o.
	// Position 1: e, position 4: o. Result: "holle"
	fmt.Println(SortVowelsByFrequency("leetcode"))
}
```
