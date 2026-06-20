# 0132 — Palindrome Partitioning Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func minCut(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #132: Palindrome Partitioning II
// https://leetcode.com/problems/palindrome-partitioning-ii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	// dp[i] = min cuts for s[0:i]
  // Alokasi slice
	dp := make([]int, n)
  // Range loop
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// isPalindrome[i][j] = s[i:j+1] is palindrome
  // Matriks 2D
	isPalindrome := make([][]bool, n)
  // Range loop
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for end := 0; end < n; end++ {
		for start := 0; start <= end; start++ {
			if s[start] == s[end] && (end-start <= 2 || isPalindrome[start+1][end-1]) {
				isPalindrome[start][end] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			dp[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if isPalindrome[j+1][i] && dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	s := "aab"
	result := minCut(s)
	expected := 1

	fmt.Printf("minCut(%q) = %d\n", s, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
