# 3622 — Check Divisibility By Digit Sum And Product

## Deskripsi

**Soal:** [3622. Check Divisibility By Digit Sum And Product](https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #3622: Check Divisibility by Digit Sum and Product
// https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(99))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(23))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(10))
}

// Time: O(log n)
// Space: O(1)
func CheckDivisibilityByDigitSumAndProduct(n int) bool {
	x := n
	sum := 0
	prod := 1
	for x > 0 {
		d := x % 10
		sum += d
		prod *= d
		x /= 10
	}
	return n%(sum+prod) == 0
}
```
