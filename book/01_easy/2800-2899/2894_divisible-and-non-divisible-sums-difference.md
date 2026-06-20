# 2894 — Divisible And Non Divisible Sums Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DivisibleAndNonDivisibleSumsDifference(n int, m int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2894: Divisible and Non-divisible Sums Difference
// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: differenceOfSums
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(10, 3)) // 19
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 6))  // 15
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 1))  // -15
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: differenceOfSums
func DivisibleAndNonDivisibleSumsDifference(n int, m int) int {
	num1 := 0 // sum of numbers not divisible by m
	num2 := 0 // sum of numbers divisible by m
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			num2 += i
		} else {
			num1 += i
		}
	}
	return num1 - num2
}
```
