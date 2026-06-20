# 3144 — Minimum Substring Partition Of Equal Character Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumSubstringsInPartition(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3144: Minimum Substring Partition of Equal Character Frequency
// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubstringsInPartition(s string) int {
	n := len(s)
  // Alokasi slice
	dp := make([]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
  // Alokasi slice
		freq := make([]int, 26)
		var distinct, maxFreq int
		for j := i - 1; j >= 0; j-- {
			idx := s[j] - 'a'
			if freq[idx] == 0 {
				distinct++
			}
			freq[idx]++
			if freq[idx] > maxFreq {
				maxFreq = freq[idx]
			}

			if maxFreq*distinct == i-j {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg")) // Expected: 3
	fmt.Println(minimumSubstringsInPartition("abababaccddb")) // Expected: ?
}
```
