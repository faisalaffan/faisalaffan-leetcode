# 2911 — Minimum Changes To Make K Semi Palindromes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumChanges(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2911: Minimum Changes to Make K Semi-palindromes
// https://leetcode.com/problems/minimum-changes-to-make-k-semi-palindromes/
// Difficulty: Hard
//
// A semi-palindrome of length L has a divisor d|L (d<L) such that grouping
// characters by residue class modulo d and each group forms a palindrome.
// Precompute cost[i][j] = min changes to make s[i:j] a semi-palindrome,
// then DP[k][n] = min changes to partition string into k semi-palindromes.
// O(N^3 * sqrt(N)) time, O(N^2) space.

import (
	"fmt"
	"math"
)

func minimumChanges(s string, k int) int {
	n := len(s)

	// Precompute cost[i][j] for substring s[i:j] (exclusive j), 0 <= i < j <= n
  // Membuat matriks/slice 2D untuk DP
	cost := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range cost {
		cost[i] = make([]int, n+1)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 2; j <= n; j++ { // need at least length 2
			length := j - i
			best := math.MaxInt32

			// Try all proper divisors d of length
			for d := 1; d < length; d++ {
				if length%d != 0 {
					continue
				}
				changes := 0
				groups := d
				groupSize := length / d

				// For each residue class (group)
				for r := 0; r < groups; r++ {
					// Characters in this group: s[i+r], s[i+r+d], s[i+r+2d], ...
					// Need each group to form a palindrome
					for p := 0; p < groupSize/2; p++ {
						leftIdx := i + r + p*d
						rightIdx := i + r + (groupSize-1-p)*d
						if s[leftIdx] != s[rightIdx] {
							changes++
						}
					}
				}

				if changes < best {
					best = changes
				}
			}

			cost[i][j] = best
		}
	}

	// DP[t][i] = min changes for first i chars into t semi-palindromes
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, k+1)
	for t := range dp {
		dp[t] = make([]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dp[t] {
			dp[t][i] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for t := 1; t <= k; t++ {
		for i := 2 * t; i <= n; i++ { // each part needs at least 2 chars
			for j := 2 * (t - 1); j <= i-2; j++ { // previous split must leave >=2 chars
				if dp[t-1][j] == math.MaxInt32 || cost[j][i] == math.MaxInt32 {
					continue
				}
				val := dp[t-1][j] + cost[j][i]
				if val < dp[t][i] {
					dp[t][i] = val
				}
			}
		}
	}

	return dp[k][n]
}

func main() {
	// Example: s="abcac", k=2 => 1
	fmt.Println(minimumChanges("abcac", 2))
	// Example: s="abcdef", k=2 => 2
	fmt.Println(minimumChanges("abcdef", 2))
	// Example: s="aabbaa", k=3 => 0
	fmt.Println(minimumChanges("aabbaa", 3))
	// Single partition (but semi-palindrome needs at least 2 chars)
	fmt.Println(minimumChanges("aba", 1))
	// k = n/2
	fmt.Println(minimumChanges("ab", 1))
	fmt.Println(minimumChanges("aabb", 2))
	// All same characters
	fmt.Println(minimumChanges("aaaa", 2))
	// Longer string
	fmt.Println(minimumChanges("abcdeabcde", 2))
}
```
