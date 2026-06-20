# 2521 — Distinct Prime Factors Of Product Of Array

## Deskripsi

**Soal:** [2521. Distinct Prime Factors Of Product Of Array](https://leetcode.com/problems/distinct-prime-factors-of-product-of-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * sqrt(max))  
**Kompleksitas Ruang:** O(number of primes)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
