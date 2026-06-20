# 3441 — Minimum Cost Good Caption

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minCostGoodCaption(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3441: Minimum Cost Good Caption
// https://leetcode.com/problems/minimum-cost-good-caption/
// Difficulty: Hard
//
// DP: partition the string into segments of length ≥ 3.
// Each segment must have all chars identical after changes.
// Cost of a segment = segment length - max frequency of any char in it.
// Find min total cost.

import "fmt"

const INF = 1 << 60

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostGoodCaption(s string) int {
	n := len(s)
	if n < 3 {
		return -1
	}

	// cost[i][j] = min cost to make s[i..j] all same character
	// We precompute for segments up to length 6 (since optimal segment rarely exceeds 6).
	// In practice, we compute on the fly.
  // Matriks 2D
	cost := make([][]int, n)
  // Range loop
	for i := range cost {
		cost[i] = make([]int, n)
	}

	// For segments of length 3 to 6 (any reasonable segment length)
	for i := 0; i < n; i++ {
		freq := [26]int{}
		maxFreq := 0
		for j := i; j < n && j-i < 10; j++ {
			freq[s[j]-'a']++
			ch := freq[s[j]-'a']
			if ch > maxFreq {
				maxFreq = ch
			}
			length := j - i + 1
			if length >= 3 {
				cost[i][j] = length - maxFreq
			} else {
				cost[i][j] = INF // cannot form segment of length < 3
			}
		}
		// For longer segments, it's never better than splitting
		for j := i + 10; j < n; j++ {
			cost[i][j] = INF
		}
	}

  // Alokasi slice
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	for i := 3; i <= n; i++ {
		for j := i - 3; j >= max(0, i-10); j-- {
			if dp[j] != INF && cost[j][i-1] != INF {
				dp[i] = min(dp[i], dp[j]+cost[j][i-1])
			}
		}
	}

	if dp[n] == INF {
		return -1
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test: "abc", expected 2
	fmt.Printf("s=abc -> %d (expected 2)\n", minCostGoodCaption("abc"))

	// Test: "aaa", expected 0
	fmt.Printf("s=aaa -> %d (expected 0)\n", minCostGoodCaption("aaa"))

	// Test: "abb", expected 1
	fmt.Printf("s=abb -> %d\n", minCostGoodCaption("abb"))

	// Test: "abcdef", expected ?
	fmt.Printf("s=abcdef -> %d\n", minCostGoodCaption("abcdef"))

	// Test: "aabbcc", expected 3 (aa|bbb|ccc etc)
	fmt.Printf("s=aabbcc -> %d\n", minCostGoodCaption("aabbcc"))

	// Test: "aaabbb", expected ? (aaa bbb: cost 0)
	fmt.Printf("s=aaabbb -> %d (expected 0)\n", minCostGoodCaption("aaabbb"))
}
```
