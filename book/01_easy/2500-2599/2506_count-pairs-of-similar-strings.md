# 2506 — Count Pairs Of Similar Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func charMask(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2506: Count Pairs Of Similar Strings
// https://leetcode.com/problems/count-pairs-of-similar-strings/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(CountPairsOfSimilarStrings([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
	fmt.Println(CountPairsOfSimilarStrings([]string{"aabb", "ab", "ba"}))                    // 3
}

func charMask(s string) int {
	mask := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func CountPairsOfSimilarStrings(words []string) int {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(words); i++ {
		maskI := charMask(words[i])
		for j := i + 1; j < len(words); j++ {
			if maskI == charMask(words[j]) {
				count++
			}
		}
	}
	return count
}
```
