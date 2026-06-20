# 3622 — Check Divisibility By Digit Sum And Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CheckDivisibilityByDigitSumAndProduct(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
