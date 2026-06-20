# 0139 — Word Break

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordBreak(s string, wordDict []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #139: Word Break
// https://leetcode.com/problems/word-break/
// Difficulty: Medium

import "fmt"

func wordBreak(s string, wordDict []string) bool {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	// Test case 1
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"})) // true

	// Test case 2
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"})) // true

	// Test case 3
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

// Time: O(n^2) | Space: O(n)
```
