# 3541 — Find Most Frequent Vowel And Consonant

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isVowel(b byte) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3541: Find Most Frequent Vowel and Consonant
// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMostFrequentVowelAndConsonant("hello world"))
	fmt.Println(FindMostFrequentVowelAndConsonant("aabbccddee"))
}

// isVowel returns true if the byte is a lowercase vowel.
func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

// FindMostFrequentVowelAndConsonant returns the sum of max vowel frequency and max consonant frequency.
// Time: O(n). Space: O(1).
func FindMostFrequentVowelAndConsonant(s string) int {
  // Alokasi slice
	vowelFreq := make([]int, 26)
  // Alokasi slice
	consonantFreq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			if isVowel(ch) {
				vowelFreq[ch-'a']++
			} else {
				consonantFreq[ch-'a']++
			}
		}
	}
	maxVowel := 0
	for _, v := range vowelFreq {
		if v > maxVowel {
			maxVowel = v
		}
	}
	maxConsonant := 0
	for _, v := range consonantFreq {
		if v > maxConsonant {
			maxConsonant = v
		}
	}
	return maxVowel + maxConsonant
}
```
