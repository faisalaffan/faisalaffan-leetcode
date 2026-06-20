# 2571 — Minimum Operations To Reduce An Integer To 0

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2571: Minimum Operations to Reduce an Integer to 0
// https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minOperations(n int) int {
	const INF = math.MaxInt32
	dp := [32][2]int{}
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = [2]int{INF, INF}
	}
	dp[0][0] = 0

	for i := 0; i < 31; i++ {
		bit := (n >> i) & 1
		for carry := 0; carry < 2; carry++ {
			if dp[i][carry] == INF {
				continue
			}
			val := bit + carry
			switch val {
			case 0:
				if dp[i][carry] < dp[i+1][0] {
					dp[i+1][0] = dp[i][carry]
				}
			case 1:
				if dp[i][carry]+1 < dp[i+1][0] {
					dp[i+1][0] = dp[i][carry] + 1
				}
				if dp[i][carry]+1 < dp[i+1][1] {
					dp[i+1][1] = dp[i][carry] + 1
				}
			case 2:
				if dp[i][carry] < dp[i+1][1] {
					dp[i+1][1] = dp[i][carry]
				}
			}
		}
	}

	ans := dp[31][0]
	if dp[31][1]+1 < ans {
		ans = dp[31][1] + 1
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations(3))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minOperations(6))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minOperations(7))
	// Expected: 2
}
```
