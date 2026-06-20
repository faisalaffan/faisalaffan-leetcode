# 1216 — Valid Palindrome Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func isValidPalindrome(s string, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1216: Valid Palindrome III
// https://leetcode.com/problems/valid-palindrome-iii/
// Difficulty: Hard [Paid]
//
// Given a string s and an integer k, determine if s can be transformed into
// a palindrome by removing at most k characters.
//
// Equivalently, check if the length of the longest palindromic subsequence
// (LPS) of s is at least len(s) - k.

import "fmt"

func main() {
	// Example 1
	fmt.Println(isValidPalindrome("abcdeca", 2)) // true (remove 'd','e' => "abca" -> "abc"+"cba"? No, remove 'd','e' => "abcca": "abcba" hmm let's see: "abcdeca", LPS length = 5 -> len-k = 7-2 = 5, so yes)

	// Example 2
	fmt.Println(isValidPalindrome("abbababa", 1)) // true

	// Simple palindrome
	fmt.Println(isValidPalindrome("aba", 0)) // true

	// Need more removals than k
	fmt.Println(isValidPalindrome("abc", 0)) // false

	// Single character (always palindrome)
	fmt.Println(isValidPalindrome("a", 0)) // true
	fmt.Println(isValidPalindrome("a", 1)) // true

	// Remove all but 1
	fmt.Println(isValidPalindrome("abcdef", 5)) // true (remove 5 -> 1 char left = palindrome)
	fmt.Println(isValidPalindrome("abcdef", 4)) // false (need LPS=2, but actual LPS=1)
}

// isValidPalindrome returns true if we can make s a palindrome by removing
// at most k characters.
//
// The length of the longest palindromic subsequence (LPS) of s can be found
// by computing the Longest Common Subsequence (LCS) between s and its reverse.
// If len(s) - LPS <= k, we can achieve a palindrome.
func isValidPalindrome(s string, k int) bool {
	n := len(s)
	lps := longestPalindromicSubsequence(s)
	return n-lps <= k
}

// longestPalindromicSubsequence returns the length of the longest palindromic
// subsequence in s using LCS(s, reverse(s)).
func longestPalindromicSubsequence(s string) int {
	n := len(s)
	// dp[i][j] = LCS of s[0..i-1] and rev[0..j-1]
	// We only need two rows.
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, 2)
	dp[0] = make([]int, n+1)
	dp[1] = make([]int, n+1)

	for i := 1; i <= n; i++ {
		cur := i % 2
		prev := 1 - cur
		for j := 1; j <= n; j++ {
			if s[i-1] == s[n-j] { // reverse(s)[j-1] == s[n-j]
				dp[cur][j] = dp[prev][j-1] + 1
			} else {
				if dp[prev][j] > dp[cur][j-1] {
					dp[cur][j] = dp[prev][j]
				} else {
					dp[cur][j] = dp[cur][j-1]
				}
			}
		}
	}

	return dp[n%2][n]
}
```
