# 3233 — Find The Count Of Numbers Which Are Not Special

## Deskripsi

**Soal:** [3233. Find The Count Of Numbers Which Are Not Special](https://leetcode.com/problems/find-the-count-of-numbers-which-are-not-special/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(r) log log r)  
**Kompleksitas Ruang:** O(sqrt(r))

**Algoritma:** —

**Fungsi Solusi:** `func nonSpecialCount(l int, r int) int`

## Solusi Go

```go
package main

// LeetCode #3233: Find the Count of Numbers Which Are Not Special
// https://leetcode.com/problems/find-the-count-of-numbers-which-are-not-special/
// Difficulty: Medium
// Time: O(sqrt(r) log log r) | Space: O(sqrt(r))

import "fmt"

func nonSpecialCount(l int, r int) int {
	limit := 31623 // sqrt(10^9) approx

  // Membuat slice untuk menyimpan hasil
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	special := 0
	for i := 2; i*i <= r; i++ {
		if isPrime[i] {
			sq := i * i
			if sq >= l && sq <= r {
				special++
			}
		}
	}

	return r - l + 1 - special
}

func main() {
	fmt.Println(nonSpecialCount(5, 7))     // Expected: 3
	fmt.Println(nonSpecialCount(4, 16))    // Expected: 11
}
```
