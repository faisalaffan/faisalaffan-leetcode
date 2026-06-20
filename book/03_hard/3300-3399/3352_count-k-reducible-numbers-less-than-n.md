# 3352 — Count K Reducible Numbers Less Than N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countKReducibleNumbers(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3352: Count K-Reducible Numbers Less Than N
// https://leetcode.com/problems/count-k-reducible-numbers-less-than-n/
// Difficulty: Hard
//
// Key insight: reducibility depends only on popcount. Use combinatorics to
// count numbers < s with each popcount. Precompute steps to reduce each
// popcount value to 1.

import "fmt"

func main() {
	// Example: "111", k=1 -> 3
	fmt.Println(countKReducibleNumbers("111", 1))
	// "100", k=1 -> 2 (numbers 1 and 2)
	fmt.Println(countKReducibleNumbers("100", 1))
	// "1", k=1 -> 0
	fmt.Println(countKReducibleNumbers("1", 1))
	// "10", k=2 -> 1 (only 1)
	fmt.Println(countKReducibleNumbers("10", 2))
	// "111", k=2 -> 5
	fmt.Println(countKReducibleNumbers("111", 2))
	// Edge: k=0, only number 1 qualifies (needs 0 steps)
	fmt.Println(countKReducibleNumbers("1000", 0))
}

const MOD = 1000000007

func countKReducibleNumbers(s string, k int) int {
	n := len(s)

	// stepsToReduceOne[v] = steps for VALUE v to reach 1 via x -> bitCount(x)
  // Alokasi slice integer
	stepsToReduceOne := make([]int, n+1)
	for i := 2; i <= n; i++ {
		stepsToReduceOne[i] = 1 + stepsToReduceOne[bitCount(i)]
	}

	// totalSteps[p] = total steps for a number with popcount p to reach 1
	// = 1 (first step: N -> p) + stepsToReduceOne[p] (for p > 1), and 0 for p = 1
  // Alokasi slice integer
	totalSteps := make([]int, n+1)
	for i := 2; i <= n; i++ {
		totalSteps[i] = 1 + stepsToReduceOne[i]
	}

	// Precompute combinations C[i][j]
  // Membuat matriks/slice 2D untuk DP
	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}

	comb := func(nn, kk int) int {
		if kk < 0 || kk > nn {
			return 0
		}
		return C[nn][kk]
	}

	// Count numbers with each popcount
  // Alokasi slice integer
	cnt := make([]int, n+1)

	// Part 1: Numbers with the same length as s but numerically smaller
	onesSoFar := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			remaining := n - i - 1
			for add := 0; add <= remaining; add++ {
				p := onesSoFar + add
				cnt[p] = (cnt[p] + comb(remaining, add)) % MOD
			}
			onesSoFar++
		}
	}

	// Part 2: Numbers with shorter lengths (1 to n-1 bits)
	// First bit is always 1 (no leading zeros)
	for length := 1; length < n; length++ {
		for p := 1; p <= length; p++ {
			// choose p-1 ones from remaining length-1 positions
			cnt[p] = (cnt[p] + comb(length-1, p-1)) % MOD
		}
	}

	// Sum qualified numbers
	ans := 0
	for p := 1; p <= n; p++ {
		if totalSteps[p] <= k {
			ans = (ans + cnt[p]) % MOD
		}
	}

	return ans
}

func bitCount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}
```
