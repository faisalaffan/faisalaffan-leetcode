# 3504 — Longest Palindrome After Substring Concatenation Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func longestPalindrome(s string, t string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3504: Longest Palindrome After Substring Concatenation II
// https://leetcode.com/problems/longest-palindrome-after-substring-concatenation-ii/
// Difficulty: Hard
//
// Given strings s and t, find the longest palindrome that can be formed by
// concatenating a substring from s (possibly empty) followed by a substring
// from t (possibly empty).
//
// Approach: Precompute longest palindromic substrings in s and t. Then find
// matching wings between reversed t and s using DP. Combine wings with
// internal palindromes.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestPalindrome("abc", "def"))
	// Example 2
	fmt.Println(longestPalindrome("a", "a"))
	// Example 3
	fmt.Println(longestPalindrome("ab", "ba"))
	// Edge: one string empty
	fmt.Println(longestPalindrome("aba", ""))
	// Edge: all matching
	fmt.Println(longestPalindrome("abc", "cba"))
}

func longestPalindrome(s string, t string) int {
	ans := 0
	m, n := len(s), len(t)

	// Precompute longest palindromic substrings in s
	// palS[i] = longest palindrome starting at or after i in s
  // Alokasi slice
	palS := make([]int, m+1)
	for i := 0; i < m; i++ {
		// odd length
		l, r := i, i
		for l >= 0 && r < m && s[l] == s[r] {
			if r-l+1 > palS[l] {
				palS[l] = r - l + 1
			}
			l--
			r++
		}
		// even length
		l, r = i, i+1
		for l >= 0 && r < m && s[l] == s[r] {
			if r-l+1 > palS[l] {
				palS[l] = r - l + 1
			}
			l--
			r++
		}
	}
	// Propagate max forward
	for i := m - 1; i >= 0; i-- {
		if palS[i+1] > palS[i] {
			palS[i] = palS[i+1]
		}
		if palS[i] > ans {
			ans = palS[i]
		}
	}

	// Precompute longest palindromes in t (ending at or before j)
  // Alokasi slice
	palT := make([]int, n+1)
	for j := 0; j < n; j++ {
		// odd length
		l, r := j, j
		for l >= 0 && r < n && t[l] == t[r] {
			if r-l+1 > palT[r] {
				palT[r] = r - l + 1
			}
			l--
			r++
		}
		// even length
		l, r = j, j+1
		for l >= 0 && r < n && t[l] == t[r] {
			if r-l+1 > palT[r] {
				palT[r] = r - l + 1
			}
			l--
			r++
		}
	}
	for j := 0; j < n; j++ {
		if palT[j] > ans {
			ans = palT[j]
		}
	}

	// DP for matching wings between s and reversed t
	// dp[i][j] = length of matching suffix between s[0..i-1] and revT[0..j-1]
	revT := reverse(t)
  // Matriks 2D
	dp := make([][]int, m+2)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == revT[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > 0 {
				// We have dp[i][j] matching characters
				// Try extending with palindrome from s remainder
				wingLen := dp[i][j] * 2
				if wingLen > ans {
					ans = wingLen
				}

				// Try extending with palindrome starting at s[i]
				if i < m && palS[i] > 0 {
					total := wingLen + palS[i]
					if total > ans {
						ans = total
					}
				}

				// Try extending with palindrome ending at t[j] (from actual t position)
				tIdx := n - j
				if tIdx >= 0 && tIdx < n && palT[tIdx] > 0 {
					total := wingLen + palT[tIdx]
					if total > ans {
						ans = total
					}
				}
			}
		}
	}

	return ans
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```
