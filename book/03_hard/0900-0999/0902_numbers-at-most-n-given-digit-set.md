# 0902 — Numbers At Most N Given Digit Set

## Deskripsi

**Soal:** [0902. Numbers At Most N Given Digit Set](https://leetcode.com/problems/numbers-at-most-n-given-digit-set/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func atMostNGivenDigitSet(digits []string, n int) int`

## Solusi Go

```go
package main

// LeetCode #902: Numbers At Most N Given Digit Set
// https://leetcode.com/problems/numbers-at-most-n-given-digit-set/
// Difficulty: Hard
//
// Combinatorics approach:
// 1. Count all numbers with fewer digits than n (all combinations).
// 2. Count numbers with same number of digits as n using digit DP:
//    For each position, try digits < current digit (rest can be anything),
//    if a digit == current digit, continue to next position.
// 3. If we matched all digits of n, add 1 for n itself.

import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	k := len(digits)

	// 1. Count numbers with fewer digits.
	result := 0
	for i := 1; i < m; i++ {
		result += pow(k, i)
	}

	// 2. Count same-length numbers.
	for i := 0; i < m; i++ {
		digitN := s[i]
		prefixMatch := false
		for _, d := range digits {
			if d[0] < digitN {
				result += pow(k, m-i-1)
			} else if d[0] == digitN {
				prefixMatch = true
				break
			}
		}
		if !prefixMatch {
			return result
		}
	}

	// 3. Exact match for n itself (we matched all positions).
	return result + 1
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	// Example 1: digits=["1","3","5","7"], n=100 -> 20
	fmt.Println("Test 1:", atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100)) // 20

	// Example 2: digits=["1","4","9"], n=1000000000 -> 29523
	fmt.Println("Test 2:", atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000)) // 29523

	// Example 3: digits=["7"], n=8 -> 1
	fmt.Println("Test 3:", atMostNGivenDigitSet([]string{"7"}, 8)) // 1

	// Edge: digits=["0"], n=0 -> 1
	fmt.Println("Test 4:", atMostNGivenDigitSet([]string{"0"}, 0)) // 1

	// Edge: digits=["1"], n=1 -> 1
	fmt.Println("Test 5:", atMostNGivenDigitSet([]string{"1"}, 1)) // 1

	// Edge: digits=["1","2","3","4","5","6","7","8","9"], n=10 -> 9 (single digit numbers only)
	fmt.Println("Test 6:", atMostNGivenDigitSet([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 10)) // 9
}
```
