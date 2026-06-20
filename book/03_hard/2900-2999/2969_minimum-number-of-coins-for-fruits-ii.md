# 2969 — Minimum Number Of Coins For Fruits Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCoins(prices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2969: Minimum Number of Coins for Fruits II
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits-ii/
// Difficulty: Hard
//
// You have n types of fruits, each with a price. When you buy fruit i,
// you get fruits i+1, i+2, ..., 2*i+1 for free (or up to n-1).
// Find minimum coins to acquire all fruits.
//
// DP from right to left: dp[i] = min cost to acquire fruits i..n-1.
// dp[i] = prices[i] + min(dp[j]) for j in [i+1, 2*i+2].
// Use monotonic deque to maintain min dp[j] in the range.

import "fmt"

func minimumCoins(prices []int) int {
	n := len(prices)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = int(1e9)
	}
	dp[n] = 0

  // Alokasi slice integer
	dq := make([]int, 0, n)
	dq = append(dq, n)

	for i := n - 1; i >= 0; i-- {
		// Remove indices that are out of range [i+1, 2*i+2]
		// Since we process right to left, we remove indices > 2*i+2
		for len(dq) > 0 && dq[0] > 2*i+2 {
			dq = dq[1:]
		}
		dp[i] = prices[i] + dp[dq[0]]
		// Maintain increasing order of dp values
		for len(dq) > 0 && dp[dq[len(dq)-1]] >= dp[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	return dp[0]
}

func main() {
	// Example
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))

	// Edge cases
	fmt.Println(minimumCoins([]int{5}))
	fmt.Println(minimumCoins([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
```
