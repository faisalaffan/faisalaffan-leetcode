# 3802 — Number Of Ways To Paint Sheets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfWays(n int, limit []int) int
```

> **💡 Hint:** Combinatorial DP. dp[i][c] = ways to paint i sheets

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3802: Number of Ways to Paint Sheets [Paid]
// https://leetcode.com/problems/number-of-ways-to-paint-sheets/
// Difficulty: Hard
//
// Count ways to paint n sheets with given color limits per color.
//
// Approach: Combinatorial DP. dp[i][c] = ways to paint i sheets
// with first c colors satisfying limits.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfWays(3, []int{2, 2}))
	// Example 2
	fmt.Println(numberOfWays(5, []int{3, 2, 1}))
	// Edge: single color
	fmt.Println(numberOfWays(2, []int{5}))
	// Edge: impossible
	fmt.Println(numberOfWays(3, []int{1, 1}))
}

func numberOfWays(n int, limit []int) int {
	const mod = 1000000007
	m := len(limit)

	// dp[i][j] = ways to paint i sheets using first j colors
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for j := 1; j <= m; j++ {
		maxUse := limit[j-1]
		for i := 0; i <= n; i++ {
			// Use k sheets of color j
			for k := 0; k <= maxUse && k <= i; k++ {
				dp[i][j] = (dp[i][j] + dp[i-k][j-1]) % mod
			}
		}
	}

	return dp[n][m]
}
```
