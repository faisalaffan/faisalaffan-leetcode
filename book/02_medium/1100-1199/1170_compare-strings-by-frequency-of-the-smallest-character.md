# 1170 — Compare Strings By Frequency Of The Smallest Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func f(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O((n + m) * L) where L = average string length  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1170: Compare Strings by Frequency of the Smallest Character
// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/
// Difficulty: Medium

// f(s) = frequency of smallest character in s.
// For each query word, count words in words[] with f(w) > f(query).

// Time: O((n + m) * L) where L = average string length
// Space: O(n)

func f(s string) int {
	minChar := s[0]
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
			count = 1
		} else if s[i] == minChar {
			count++
		}
	}
	return count
}

func numSmallerByFrequency(queries []string, words []string) []int {
  // Alokasi slice integer
	wordFreqs := make([]int, len(words))
	for i, w := range words {
		wordFreqs[i] = f(w)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(wordFreqs)

  // Alokasi slice integer
	result := make([]int, len(queries))
	for i, q := range queries {
		qf := f(q)
		// Binary search for first wordFreq > qf
		idx := sort.Search(len(wordFreqs), func(j int) bool {
			return wordFreqs[j] > qf
		})
		result[i] = len(wordFreqs) - idx
	}
	return result
}

func main() {
	fmt.Printf("%v (expected: [1])\n", numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Printf("%v (expected: [1 2])\n", numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
```
