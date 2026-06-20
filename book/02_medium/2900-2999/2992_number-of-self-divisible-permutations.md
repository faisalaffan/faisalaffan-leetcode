# 2992 — Number Of Self Divisible Permutations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSelfDivisiblePermutations(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, GCD / Matematika, Bitmask

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2992: Number of Self-Divisible Permutations
// https://leetcode.com/problems/number-of-self-divisible-permutations/
// Difficulty: Medium

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func numberOfSelfDivisiblePermutations(n int) int {
	maskCount := 1 << n
  // Alokasi slice integer
	dp := make([]int, maskCount)
	dp[0] = 1

	// position (1-indexed) = number of set bits in the mask
	// Precompute popcount for each mask
  // Alokasi slice integer
	popcount := make([]int, maskCount)
	for mask := 1; mask < maskCount; mask++ {
		popcount[mask] = popcount[mask>>1] + (mask & 1)
	}

	for mask := 1; mask < maskCount; mask++ {
		pos := popcount[mask] // 1-indexed position
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 && gcd(pos, j+1) == 1 {
				dp[mask] += dp[mask^(1<<j)]
			}
		}
	}

	return dp[maskCount-1]
}

func main() {
	// Test case 1: n = 1 -> 1 (only [1])
	fmt.Println(numberOfSelfDivisiblePermutations(1)) // 1

	// Test case 2: n = 2 -> 1 (only [2,1])
	fmt.Println(numberOfSelfDivisiblePermutations(2)) // 1

	// Test case 3: n = 3 -> 3 ([1,3,2], [3,1,2], [2,3,1])
	fmt.Println(numberOfSelfDivisiblePermutations(3)) // 3

	// Test case 4: n = 4
	fmt.Println(numberOfSelfDivisiblePermutations(4)) // ?

	// Test case 5: n = 5
	fmt.Println(numberOfSelfDivisiblePermutations(5)) // ?
}

// Time: O(n * 2^n) | Space: O(2^n)
```
