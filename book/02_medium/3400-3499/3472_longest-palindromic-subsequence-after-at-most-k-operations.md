# 3472 — Longest Palindromic Subsequence After At Most K Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestPalindromicSubsequenceAfterAtMostKOperations(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3472: Longest Palindromic Subsequence After at Most K Operations
// https://leetcode.com/problems/longest-palindromic-subsequence-after-at-most-k-operations/
// Difficulty: Medium
// Complexity: O(n^2 * k) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", LongestPalindromicSubsequenceAfterAtMostKOperations("abcde", 2))
	// Test case 2
	fmt.Println("Test 2:", LongestPalindromicSubsequenceAfterAtMostKOperations("abac", 1))
	// Test case 3
	fmt.Println("Test 3:", LongestPalindromicSubsequenceAfterAtMostKOperations("a", 0))
}

func LongestPalindromicSubsequenceAfterAtMostKOperations(s string, k int) int {
	n := len(s)
	// dp[i][j][t] = longest palindromic subsequence in s[i..j] using at most t operations
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for t := 0; t <= k; t++ {
		for i := 0; i < n; i++ {
			dp[i][i][t] = 1
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			for t := 0; t <= k; t++ {
				// skip left
				if dp[i+1][j][t] > dp[i][j][t] {
					dp[i][j][t] = dp[i+1][j][t]
				}
				// skip right
				if dp[i][j-1][t] > dp[i][j][t] {
					dp[i][j][t] = dp[i][j-1][t]
				}
				// match
				cost := diff(s[i], s[j])
				if cost <= t {
					base := 2
					if i+1 <= j-1 {
						base += dp[i+1][j-1][t-cost]
					}
					if base > dp[i][j][t] {
						dp[i][j][t] = base
					}
				}
			}
		}
	}

	return dp[0][n-1][k]
}

func diff(a, b byte) int {
	d := int(a) - int(b)
	if d < 0 {
		d = -d
	}
	return d
}
```
