# 3639 — Minimum Time To Activate String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumTimeToActivateString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3639: Minimum Time to Activate String
// https://leetcode.com/problems/minimum-time-to-activate-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minimumTimeToActivateString(s string) int {
	n := len(s)
	// dp[i] = min time to activate prefix of length i
  // Alokasi slice
	dp := make([]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = i // worst case: type each character
	}

	for i := 1; i <= n; i++ {
		// Option 1: type the current character (1 sec)
		if dp[i] > dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
		}

		// Option 2: try to activate from a matching prefix
		// Look for a substring in already activated prefix that matches s[i-1:]
		for j := 1; j < i; j++ {
			k := 0
			for i-1+k < n && j-1+k < i-1 && s[i-1+k] == s[j-1+k] {
				k++
			}
			if k > 0 {
				cost := dp[i-1] + 1 // 1 second to activate the matching substring
				if dp[i+k-1] > cost {
					dp[i+k-1] = cost
				}
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(minimumTimeToActivateString("abcabc"))
	fmt.Println(minimumTimeToActivateString("aaaa"))
	fmt.Println(minimumTimeToActivateString("abacaba"))
}
```
