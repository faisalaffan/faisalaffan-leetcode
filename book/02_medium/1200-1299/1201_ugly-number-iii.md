# 1201 — Ugly Number Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func nthUglyNumber(n int, a int, b int, c int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, GCD / Matematika

**Kompleksitas Waktu:** O(log(maxVal)) ~ O(log(2*10^9))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1201: Ugly Number III
// https://leetcode.com/problems/ugly-number-iii/
// Difficulty: Medium

// Find the nth number divisible by a, b, or c.
// Binary search on answer + inclusion-exclusion principle.

// Time: O(log(maxVal)) ~ O(log(2*10^9))
// Space: O(1)

func nthUglyNumber(n int, a int, b int, c int) int {
	// LCM helpers
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	lcm := func(x, y int) int {
		return x / gcd(x, y) * y
	}

	ab := lcm(a, b)
	bc := lcm(b, c)
	ac := lcm(a, c)
	abc := lcm(a, lcm(b, c))

	countUgly := func(num int) int {
		return num/a + num/b + num/c - num/ab - num/ac - num/bc + num/abc
	}

	lo, hi := 1, 2000000000
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countUgly(mid) >= n {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Printf("%d (expected: 4)\n", nthUglyNumber(3, 2, 3, 5))
	fmt.Printf("%d (expected: 6)\n", nthUglyNumber(4, 2, 3, 4))
	fmt.Printf("%d (expected: 10)\n", nthUglyNumber(5, 2, 11, 13))
}
```
