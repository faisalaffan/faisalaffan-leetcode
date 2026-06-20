# 1269 — Number Of Ways To Stay In The Same Place After Some Steps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numWays(steps int, arrLen int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1269: Number of Ways to Stay in the Same Place After Some Steps
// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1269. Number of Ways to Stay in the Same Place After Some Steps")
	fmt.Println("steps=3, arrLen=3:", numWays(3, 3), "(expected 4)")
	fmt.Println("steps=2, arrLen=4:", numWays(2, 4), "(expected 2)")
	fmt.Println("steps=4, arrLen=2:", numWays(4, 2), "(expected 8)")
}

func numWays(steps int, arrLen int) int {
	const MOD = 1000000007

	// Max reachable position is min(steps, arrLen-1).
	maxPos := steps
	if arrLen-1 < maxPos {
		maxPos = arrLen - 1
	}

  // Alokasi slice integer
	dp := make([]int, maxPos+1)
	dp[0] = 1

	for s := 1; s <= steps; s++ {
  // Alokasi slice integer
		ndp := make([]int, maxPos+1)
		for pos := 0; pos <= maxPos; pos++ {
			ways := dp[pos] // stay
			if pos > 0 {
				ways = (ways + dp[pos-1]) % MOD // move right (from pos-1)
			}
			if pos < maxPos {
				ways = (ways + dp[pos+1]) % MOD // move left (from pos+1)
			}
			ndp[pos] = ways
		}
		dp = ndp
	}

	return dp[0]
}
```
