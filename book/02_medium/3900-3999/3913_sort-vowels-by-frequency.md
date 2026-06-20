# 3913 — Sort Vowels By Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func isVowel(ch byte) bool
```

> **💡 Hint:** Count vowel frequencies, sort vowels by freq desc (tie: alpha),

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
  // Custom sort dengan comparator
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
