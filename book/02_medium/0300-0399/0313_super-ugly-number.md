# 0313 — Super Ugly Number

## Deskripsi

**Soal:** [0313. Super Ugly Number](https://leetcode.com/problems/super-ugly-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * len(primes)), Space: O(n + len(primes))  
**Kompleksitas Ruang:** O(n + len(primes))

**Algoritma:** —

**Fungsi Solusi:** `func nthSuperUglyNumber(n int, primes []int) int`

## Solusi Go

```go
package main

// LeetCode #313: Super Ugly Number
// https://leetcode.com/problems/super-ugly-number/
// Difficulty: Medium
// Time: O(n * len(primes)), Space: O(n + len(primes))

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
  // Membuat slice untuk menyimpan hasil
	ugly := make([]int, n)
	ugly[0] = 1

  // Membuat slice untuk menyimpan hasil
	pointers := make([]int, len(primes))
  // Membuat slice untuk menyimpan hasil
	values := make([]int, len(primes))
	for i, p := range primes {
		values[i] = p
	}

	for i := 1; i < n; i++ {
		minVal := values[0]
		for _, v := range values {
			if v < minVal {
				minVal = v
			}
		}
		ugly[i] = minVal

		for j := range values {
			if values[j] == minVal {
				pointers[j]++
				values[j] = ugly[pointers[j]] * primes[j]
			}
		}
	}

	return ugly[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}))
	fmt.Println(nthSuperUglyNumber(6, []int{2, 3, 5}))
}
```
