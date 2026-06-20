# 1930 — Unique Length 3 Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountPalindromicSubsequence(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * 26) = O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1930: Unique Length-3 Palindromic Subsequences
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountPalindromicSubsequence("aabca"))
	fmt.Println(CountPalindromicSubsequence("adc"))
	fmt.Println(CountPalindromicSubsequence("bbcbaba"))
}

// Time: O(n * 26) = O(n), Space: O(1)
func CountPalindromicSubsequence(s string) int {
	// For each character, find first and last occurrence
  // Alokasi slice
	first := make([]int, 26)
  // Alokasi slice
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	count := 0
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c]-first[c] > 1 {
			// Count unique characters between first and last occurrence
			seen := make([]bool, 26)
			for i := first[c] + 1; i < last[c]; i++ {
				seen[s[i]-'a'] = true
			}
			for _, v := range seen {
				if v {
					count++
				}
			}
		}
	}
	return count
}
```
