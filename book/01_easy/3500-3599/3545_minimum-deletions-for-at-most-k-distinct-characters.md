# 3545 — Minimum Deletions For At Most K Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumDeletionsForAtMostKDistinctCharacters(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3545: Minimum Deletions for At Most K Distinct Characters
// https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("aabbbcc", 2))
	fmt.Println(MinimumDeletionsForAtMostKDistinctCharacters("abcde", 2))
}

// MinimumDeletionsForAtMostKDistinctCharacters returns min deletions so the string has at most k distinct characters.
// Time: O(n log n). Space: O(1).
func MinimumDeletionsForAtMostKDistinctCharacters(s string, k int) int {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
  // Custom sort dengan comparator
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	// Count distinct characters
	distinct := 0
	for _, f := range freq {
		if f > 0 {
			distinct++
		}
	}
	if distinct <= k {
		return 0
	}

	// Delete the least frequent characters (from the end of sorted freq)
	deletions := 0
	for i := len(freq) - 1; i >= 0 && distinct > k; i-- {
		if freq[i] > 0 {
			deletions += freq[i]
			distinct--
		}
	}
	return deletions
}
```
