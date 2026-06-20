# 1745 — Palindrome Partitioning Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkPartitioning(s string) bool
```

> **💡 Hint:** DP palindrome table + split check.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1745: Palindrome Partitioning IV
// https://leetcode.com/problems/palindrome-partitioning-iv/
// Difficulty: Hard
//
// Approach: DP palindrome table + split check.
// 1. Precompute isPal[i][j] = true if s[i..j] is palindrome.
// 2. Check pairs of split points (i, j) such that:
//    isPal[0][i-1] && isPal[i][j-1] && isPal[j][n-1] are all true.

import "fmt"

func checkPartitioning(s string) bool {
	n := len(s)
	// dp[i][j] = s[i..j] is palindrome
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]bool, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// All single chars are palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// Two chars
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	// Longer substrings
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}

	// Try all split points: first partition ends at i-1, second ends at j-1
	for i := 1; i < n-1; i++ {
		if !dp[0][i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if dp[i][j-1] && dp[j][n-1] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Example test cases
	fmt.Println("\"abcbdd\" →", checkPartitioning("abcbdd")) // Expected: true
	fmt.Println("\"abc\" →", checkPartitioning("abc"))       // Expected: false
	fmt.Println("\"aabb\" →", checkPartitioning("aabb"))     // Expected: true (aa|bb|or aa|b|b etc)
	fmt.Println("\"abca\" →", checkPartitioning("abca"))     // Expected: false
	fmt.Println("\"aaaa\" →", checkPartitioning("aaaa"))     // Expected: true
}
```
