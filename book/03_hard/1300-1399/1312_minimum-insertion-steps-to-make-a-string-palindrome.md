# 1312 — Minimum Insertion Steps To Make A String Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func minInsertions(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1312: Minimum Insertion Steps to Make a String Palindrome
// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
// Difficulty: Hard
//
// Approach: DP on intervals.
// dp[i][j] = minimum insertions needed to make s[i..j] a palindrome.
//   If s[i] == s[j]: dp[i][j] = dp[i+1][j-1] (0 for length < 2)
//   Else: dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
// Answer = dp[0][n-1].

import "fmt"

func minInsertions(s string) int {
	n := len(s)
  // Matriks 2D
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if i+1 <= j-1 {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
			}
		}
	}

  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minInsertions("zzazz"))     // 0
	fmt.Println(minInsertions("mbadm"))     // 2
	fmt.Println(minInsertions("leetcode"))  // 5
	fmt.Println(minInsertions("g"))         // 0
	fmt.Println(minInsertions("no"))        // 1
	fmt.Println(minInsertions(""))          // 0
}
```
