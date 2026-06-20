# 2707 — Extra Characters In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minExtraChar(s string, dictionary []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2707: Extra Characters in a String
// https://leetcode.com/problems/extra-characters-in-a-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minExtraChar(s string, dictionary []string) int {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range dictionary {
		wordSet[w] = true
	}

	n := len(s)
  // Alokasi slice
	dp := make([]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		// Try to match word ending at i-1
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if wordSet[s[j:i]] {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	// Expected: 1 ('s' is extra)

	// Test case 2
	fmt.Println("Test 2:", minExtraChar("sayhelloworld", []string{"hello", "world"}))
	// Expected: 3 ('say' is extra)

	// Test case 3
	fmt.Println("Test 3:", minExtraChar("abcd", []string{"a", "b", "c", "d"}))
	// Expected: 0
}
```
