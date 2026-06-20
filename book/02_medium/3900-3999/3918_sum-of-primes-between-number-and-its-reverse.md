# 3918 — Sum Of Primes Between Number And Its Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SumOfPrimesBetweenNumberAndItsReverse(n int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N log log N)  |  **Ruang:** O(N) where N = max(n, rev(n)) <= 1000


## 💻 Solusi Go

```go
package main

// LeetCode #3918: Sum of Primes Between Number and Its Reverse
// https://leetcode.com/problems/sum-of-primes-between-number-and-its-reverse/
// Difficulty: Medium
// Time: O(N log log N) | Space: O(N) where N = max(n, rev(n)) <= 1000
// Approach: Sieve primes up to hi = max(n, rev(n)), sum primes in [lo, hi].

import "fmt"

func SumOfPrimesBetweenNumberAndItsReverse(n int) int64 {
	rev := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		rev = rev*10 + tmp%10
	}

	lo, hi := n, rev
	if lo > hi {
		lo, hi = hi, lo
	}

	isPrime := make([]bool, hi+1)
	for i := 2; i <= hi; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= hi; i++ {
		if isPrime[i] {
			for j := i * i; j <= hi; j += i {
				isPrime[j] = false
			}
		}
	}

	var sum int64 = 0
	for i := lo; i <= hi; i++ {
		if isPrime[i] {
			sum += int64(i)
		}
	}
	return sum
}

func main() {
	// Example 1
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(13)) // Expected: 132

	// Example 2
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(10)) // Expected: 17

	// Example 3
	fmt.Println(SumOfPrimesBetweenNumberAndItsReverse(8)) // Expected: 0
}
```
