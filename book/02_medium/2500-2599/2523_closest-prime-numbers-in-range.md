# 2523 — Closest Prime Numbers In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func closestPrimes(left int, right int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(right log log right)  |  **Ruang:** O(right)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2523: Closest Prime Numbers in Range
// https://leetcode.com/problems/closest-prime-numbers-in-range/
// Difficulty: Medium
// Time: O(right log log right) | Space: O(right)
// Sieve of Eratosthenes, find adjacent primes with min difference.

import "fmt"

func main() {
	fmt.Println(closestPrimes(10, 19)) // [11, 13]
	fmt.Println(closestPrimes(4, 6))   // [-1, -1]
}

func closestPrimes(left int, right int) []int {
	if right < 2 {
		return []int{-1, -1}
	}

	isPrime := make([]bool, right+1)
	for i := 2; i <= right; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= right; i++ {
		if isPrime[i] {
			for j := i * i; j <= right; j += i {
				isPrime[j] = false
			}
		}
	}

	prev := -1
	minDiff := right + 1
	ans := []int{-1, -1}

	for i := left; i <= right; i++ {
		if isPrime[i] {
			if prev != -1 && i-prev < minDiff {
				minDiff = i - prev
				ans = []int{prev, i}
			}
			prev = i
		}
	}
	return ans
}
```
