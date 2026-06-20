# 0960 — Delete Columns To Make Sorted Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minDeletionSize(A []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #960: Delete Columns to Make Sorted III
// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/
// Difficulty: Hard

import "fmt"

func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	m, n := len(A), len(A[0])

	// dp[j] = longest increasing subsequence ending at column j
  // Alokasi slice
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}

	for j := 0; j < n; j++ {
		for k := 0; k < j; k++ {
			// Check if we can place column j after column k
			ok := true
			for i := 0; i < m; i++ {
				if A[i][k] > A[i][j] {
					ok = false
					break
				}
			}
			if ok {
				dp[j] = max(dp[j], dp[k]+1)
			}
		}
	}

	// max columns we can keep
	maxKeep := 0
	for _, v := range dp {
		maxKeep = max(maxKeep, v)
	}
	// min columns to delete = total - maxKeep
	return n - maxKeep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(minDeletionSize([]string{"babca", "bbazb"}))
	// Expected: 3

	fmt.Println("Example 2:")
	fmt.Println(minDeletionSize([]string{"edcba"}))
	// Expected: 4

	fmt.Println("Example 3:")
	fmt.Println(minDeletionSize([]string{"ghi", "def", "abc"}))
	// Expected: 0
}
```
