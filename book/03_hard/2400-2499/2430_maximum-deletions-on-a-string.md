# 2430 — Maximum Deletions On A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDeletions(s string) int
```

> **💡 Hint:** DP + LCP (Longest Common Prefix).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2430: Maximum Deletions on a String
// https://leetcode.com/problems/maximum-deletions-on-a-string/
// Difficulty: Hard
//
// Given a string s. In one operation, you can delete a prefix of the current
// string if the remaining string starts with that same prefix. Find the maximum
// number of operations needed to delete the entire string.
//
// Approach: DP + LCP (Longest Common Prefix).
// dp[i] = max deletions starting at position i (suffix s[i:]).
// lcp[i][j] = longest common prefix of s[i:] and s[j:].
//
// For each i, for each possible length len (where i+len < n):
//   if lcp[i][i+len] >= len (i.e., s[i:i+len] == s[i+len:i+2*len]),
//   then dp[i] = max(dp[i], 1 + dp[i+len]).
//
// Base: dp[n] = 0 (empty string). Result: dp[0].

import "fmt"

func maxDeletions(s string) int {
	n := len(s)

	// lcp[i][j] = longest common prefix of s[i:] and s[j:]
  // Membuat matriks/slice 2D untuk DP
	lcp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = 1 + lcp[i+1][j+1]
			}
		}
	}

  // Alokasi slice integer
	dp := make([]int, n+1)
	// dp[n] = 0 by default

	for i := n - 1; i >= 0; i-- {
		dp[i] = 1 // at minimum, we can delete the whole suffix in one operation
		maxLen := (n - i) / 2
		for length := 1; length <= maxLen; length++ {
			if lcp[i][i+length] >= length {
				if 1+dp[i+length] > dp[i] {
					dp[i] = 1 + dp[i+length]
				}
			}
		}
	}

	return dp[0]
}

func main() {
	// Example 1
	fmt.Println(maxDeletions("abcabcdabc"))
	// Example 2
	fmt.Println(maxDeletions("aaabaab"))
	// Example 3
	fmt.Println(maxDeletions("aaaaa"))
	// Single char
	fmt.Println(maxDeletions("a"))
	// Two chars
	fmt.Println(maxDeletions("aa"))
	fmt.Println(maxDeletions("ab"))
	// No repeated prefix
	fmt.Println(maxDeletions("abcdef"))
}
```
