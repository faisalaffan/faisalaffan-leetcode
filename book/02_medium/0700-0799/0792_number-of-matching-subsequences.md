# 0792 — Number Of Matching Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numMatchingSubseq(s string, words []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m * L) where n = len(s), m = len(words)  |  **Ruang:** O(m)


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
  // Matriks 2D
	buckets := make([][]string, 26)
  // Range loop
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
