# 3316 — Find Maximum Removals From Source String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maxRemovals(source string, pattern string, targetIndices []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * m) Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3316: Find Maximum Removals From Source String
// https://leetcode.com/problems/find-maximum-removals-from-source-string/
// Difficulty: Medium
// Time: O(n * m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(maxRemovals("abc", "ab", []int{0, 1}))                          // 0
	fmt.Println(maxRemovals("abbaa", "aba", []int{0, 1, 2, 3, 4}))              // 2
	fmt.Println(maxRemovals("abcde", "ace", []int{1, 2, 3}))                    // 1
}

func maxRemovals(source string, pattern string, targetIndices []int) int {
	n, m := len(source), len(pattern)
	target := make([]bool, n)
	for _, idx := range targetIndices {
		target[idx] = true
	}

  // Alokasi slice
	dp := make([]int, m+1)
  // Range loop
	for i := range dp {
		dp[i] = -1_000_000_000
	}
	dp[m] = 0

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			if target[i] {
				dp[j]++
			}
			if j < m && source[i] == pattern[j] {
				if dp[j]+1 > dp[j] && dp[j+1] > dp[j] {
					if dp[j+1] > dp[j] {
						dp[j] = dp[j+1]
					}
				} else if dp[j+1] > dp[j] {
					dp[j] = dp[j+1]
				}
			}
		}
	}

	return dp[0]
}
```
