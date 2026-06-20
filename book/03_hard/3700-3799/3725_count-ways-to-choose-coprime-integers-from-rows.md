# 3725 — Count Ways To Choose Coprime Integers From Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countWays(mat [][]int) int
```

> **💡 Hint:** DP over GCD values. For each row, compute new GCDs.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming, GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3725: Count Ways to Choose Coprime Integers from Rows
// https://leetcode.com/problems/count-ways-to-choose-coprime-integers-from-rows/
// Difficulty: Hard
//
// Pick exactly one integer from each row such that GCD of all
// chosen integers is 1 (coprime). Return count modulo 1e9+7.
//
// Approach: DP over GCD values. For each row, compute new GCDs.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countWays([][]int{{2, 3}, {3, 5}}))
	// Example 2
	fmt.Println(countWays([][]int{{2, 4}, {6, 8}}))
	// Edge: single row
	fmt.Println(countWays([][]int{{1, 5}}))
	// Edge: already coprime
	fmt.Println(countWays([][]int{{1, 2}, {3, 4}}))
}

func countWays(mat [][]int) int {
	const mod = 1000000007
	m := len(mat)
	if m == 0 {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	dp := make(map[int]int)
	for _, v := range mat[0] {
		dp[v] = (dp[v] + 1) % mod
	}

	for r := 1; r < m; r++ {
  // Membuat map (HashMap) — pencarian O(1)
		ndp := make(map[int]int)
		for g, cnt := range dp {
			for _, v := range mat[r] {
				ng := gcd(g, v)
				ndp[ng] = (ndp[ng] + cnt) % mod
			}
		}
		dp = ndp
	}

	return dp[1]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
