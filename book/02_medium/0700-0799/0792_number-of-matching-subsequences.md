# 0792 — Number Of Matching Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numMatchingSubseq(s string, words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m * L) where n = len(s), m = len(words)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #792: Number of Matching Subsequences
// https://leetcode.com/problems/number-of-matching-subsequences/
// Difficulty: Medium
// Time: O(n + m * L) where n = len(s), m = len(words)
// Space: O(m)

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func numMatchingSubseq(s string, words []string) int {
  // Membuat matriks/slice 2D untuk DP
	buckets := make([][]string, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range buckets {
		buckets[i] = make([]string, 0)
	}

	for _, w := range words {
		buckets[w[0]-'a'] = append(buckets[w[0]-'a'], w)
	}

	count := 0

	for _, c := range s {
		idx := c - 'a'
		curr := buckets[idx]
		buckets[idx] = make([]string, 0)

		for _, w := range curr {
			if len(w) == 1 {
				count++
			} else {
				buckets[w[1]-'a'] = append(buckets[w[1]-'a'], w[1:])
			}
		}
	}

	return count
}
```
