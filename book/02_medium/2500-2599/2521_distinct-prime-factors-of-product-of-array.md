# 2521 — Distinct Prime Factors Of Product Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func distinctPrimeFactors(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * sqrt(max))  |  **Ruang:** O(number of primes)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2521: Distinct Prime Factors of Product of Array
// https://leetcode.com/problems/distinct-prime-factors-of-product-of-array/
// Difficulty: Medium
// Time: O(n * sqrt(max)) | Space: O(number of primes)
// Find union of prime factors of all numbers.

import "fmt"

func main() {
	fmt.Println(distinctPrimeFactors([]int{2, 4, 3, 7, 10, 6})) // 4 (2, 3, 5, 7)
	fmt.Println(distinctPrimeFactors([]int{4, 8, 16}))           // 1 (2)
}

func distinctPrimeFactors(nums []int) int {
  // HashMap: O(1) lookup
	primes := make(map[int]bool)
	for _, v := range nums {
		addPrimeFactors(v, primes)
	}
	return len(primes)
}

func addPrimeFactors(n int, set map[int]bool) {
	// Factor 2
	if n%2 == 0 {
		set[2] = true
		for n%2 == 0 {
			n /= 2
		}
	}
	// Odd factors
	for f := 3; f*f <= n; f += 2 {
		if n%f == 0 {
			set[f] = true
			for n%f == 0 {
				n /= f
			}
		}
	}
	if n > 1 {
		set[n] = true
	}
}
```
