# 0730 — Count Different Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countPalindromicSubsequences(S string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #730: Count Different Palindromic Subsequences
// https://leetcode.com/problems/count-different-palindromic-subsequences/
// Difficulty: Hard
//
// Algorithm: 3D Interval DP
// dp[i][j][k] = count of distinct palindromic subsequences in S[i:j+1]
// that start and end with character k (where k=0..3 maps to a,b,c,d)
//
// For each interval [i,j]:
//   if S[i] != S[j]:
//     dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
//   else:
//     Find next occurrence of S[i] after i, and prev occurrence before j
//     If no inner occurrence: 2 + dp[i+1][j-1]*2
//     If one inner occurrence: 1 + dp[i+1][j-1]*2
//     If multiple: dp[i+1][j-1]*2 - dp[next+1][prev-1]
//
// Since letters are only a,b,c,d (4 chars), we can use a 2D DP with careful
// handling of duplicates using next/prev arrays.

import (
	"fmt"
)

const mod = 1000000007

func countPalindromicSubsequences(S string) int {
	n := len(S)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// dp[i][j] = count of distinct palindromic subsequences in S[i:j+1]
  // Matriks 2D
	dp := make([][]int, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Precompute next and prev occurrences for each position
	// nextPos[i][c] = next position >= i with char c (or -1)
	// prevPos[i][c] = previous position <= i with char c (or -1)
	const letters = 4
  // Alokasi slice
	nextPos := make([][letters]int, n)
  // Alokasi slice
	prevPos := make([][letters]int, n)

	// Initialize nextPos from right to left
	last := [letters]int{-1, -1, -1, -1}
	for i := n - 1; i >= 0; i-- {
		last[S[i]-'a'] = i
		for c := 0; c < letters; c++ {
			nextPos[i][c] = last[c]
		}
	}

	// Initialize prevPos from left to right
	last = [letters]int{-1, -1, -1, -1}
	for i := 0; i < n; i++ {
		last[S[i]-'a'] = i
		for c := 0; c < letters; c++ {
			prevPos[i][c] = last[c]
		}
	}

	// Single character substrings
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// Process by increasing length
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			if S[i] != S[j] {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
				if dp[i][j] < 0 {
					dp[i][j] += mod
				}
				dp[i][j] %= mod
			} else {
				c := S[i] - 'a'
				next := nextPos[i+1][c]
				prev := prevPos[j-1][c]

				if next > prev || next == -1 {
					// No occurrence of S[i] inside (i+1, j-1)
					dp[i][j] = dp[i+1][j-1]*2 + 2
				} else if next == prev {
					// Exactly one occurrence inside
					dp[i][j] = dp[i+1][j-1]*2 + 1
				} else {
					// Two or more occurrences inside
					dp[i][j] = dp[i+1][j-1]*2 - dp[next+1][prev-1]
					if dp[i][j] < 0 {
						dp[i][j] += mod
					}
				}
				dp[i][j] %= mod
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	// Example from problem
	S1 := "bccb"
	result1 := countPalindromicSubsequences(S1)
	fmt.Printf("Input: %q\nOutput: %d (expected: 6)\n\n", S1, result1)

	// Test case 2
	S2 := "abcd"
	result2 := countPalindromicSubsequences(S2)
	fmt.Printf("Input: %q\nOutput: %d (expected: 4)\n\n", S2, result2)

	// Test case 3: all same characters
	S3 := "aaaa"
	result3 := countPalindromicSubsequences(S3)
	fmt.Printf("Input: %q\nOutput: %d\n\n", S3, result3)

	// Test case 4
	S4 := "a"
	result4 := countPalindromicSubsequences(S4)
	fmt.Printf("Input: %q\nOutput: %d (expected: 1)\n\n", S4, result4)

	// Test case 5
	S5 := "aa"
	result5 := countPalindromicSubsequences(S5)
	fmt.Printf("Input: %q\nOutput: %d\n\n", S5, result5)

	// Test case 6
	S6 := "ab"
	result6 := countPalindromicSubsequences(S6)
	fmt.Printf("Input: %q\nOutput: %d (expected: 2: a, b)\n\n", S6, result6)

	// Test case 7
	S7 := "bcba"
	result7 := countPalindromicSubsequences(S7)
	fmt.Printf("Input: %q\nOutput: %d\n\n", S7, result7)

	// Test case 8: longer example
	S8 := "abcdabcdabcdabcd"
	result8 := countPalindromicSubsequences(S8)
	fmt.Printf("Input: %q\nOutput: %d\n", S8, result8)
}
```
