# 0762 — Prime Number Of Set Bits In Binary Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countPrimeSetBits(left int, right int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O((right-left+1) * log n). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #762: Prime Number of Set Bits in Binary Representation
// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(6, 10))   // 4
	fmt.Println(countPrimeSetBits(10, 15))  // 5
	fmt.Println(countPrimeSetBits(1, 2))    // 1
}

// countPrimeSetBits counts numbers in [left, right] whose binary representation has a prime number of set bits.
// Time: O((right-left+1) * log n). Space: O(1).
func countPrimeSetBits(left int, right int) int {
	// Primes up to 20 (since max int is 10^6, < 2^20)
	primes := map[int]bool{2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true}
	count := 0
	for n := left; n <= right; n++ {
		bits := 0
		for x := n; x > 0; x >>= 1 {
			bits += x & 1
		}
		if primes[bits] {
			count++
		}
	}
	return count
}
```
