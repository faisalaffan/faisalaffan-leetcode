# 2901 — Longest Unequal Adjacent Groups Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2901: Longest Unequal Adjacent Groups Subsequence II
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string {
	n := len(words)
  // Alokasi slice
	dp := make([]int, n)
  // Alokasi slice
	prev := make([]int, n)
  // Range loop
	for i := range prev {
		prev[i] = -1
	}

	hamming := func(a, b string) int {
		if len(a) != len(b) {
			return -1
		}
		diff := 0
  // Linear scan O(n)
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff
	}

	bestLen := 0
	bestIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && hamming(words[j], words[i]) == 1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
		if dp[i] > bestLen {
			bestLen = dp[i]
			bestIdx = i
		}
	}

	result := make([]string, bestLen)
	for i := bestLen - 1; i >= 0; i-- {
		result[i] = words[bestIdx]
		bestIdx = prev[bestIdx]
	}

	return result
}

func main() {
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
}
```
