# 1278 — Palindrome Partitioning Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func palindromePartition(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1278: Palindrome Partitioning III
// https://leetcode.com/problems/palindrome-partitioning-iii/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1278. Palindrome Partitioning III")
	fmt.Println("\"abc\", k=2:", palindromePartition("abc", 2), "(expected 1)")
	fmt.Println("\"aabbc\", k=3:", palindromePartition("aabbc", 3), "(expected 0)")
	fmt.Println("\"leetcode\", k=8:", palindromePartition("leetcode", 8), "(expected 0)")
}

func palindromePartition(s string, k int) int {
	n := len(s)
	if k >= n {
		return 0
	}

	// cost[i][j] = min changes to make s[i:j+1] a palindrome.
  // Membuat matriks/slice 2D untuk DP
	cost := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			cost[i][j] = cost[i+1][j-1]
			if s[i] != s[j] {
				cost[i][j]++
			}
		}
	}

	// dp[i][p] = min changes for s[:i] split into p palindromes.
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k+1)
		for p := range dp[i] {
			dp[i][p] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for p := 1; p <= k && p <= i; p++ {
			if p == 1 {
				dp[i][1] = cost[0][i-1]
			} else {
				// Try all possible last partition boundaries.
				for j := p - 1; j < i; j++ {
					if dp[j][p-1] != math.MaxInt32 {
						val := dp[j][p-1] + cost[j][i-1]
						if val < dp[i][p] {
							dp[i][p] = val
						}
					}
				}
			}
		}
	}

	return dp[n][k]
}
```
