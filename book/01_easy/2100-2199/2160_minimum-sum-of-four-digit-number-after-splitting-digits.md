# 2160 — Minimum Sum Of Four Digit Number After Splitting Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumSumOfFourDigitNumberAfterSplittingDigits(num int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2160: Minimum Sum of Four Digit Number After Splitting Digits
// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(2932)) // 52
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(4009)) // 13
}

// Time: O(1), Space: O(1)
func MinimumSumOfFourDigitNumberAfterSplittingDigits(num int) int {
  // Alokasi slice
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = num % 10
		num /= 10
	}
  // Sort O(n log n)
	sort.Ints(digits)
	// Smallest sum: smallest digit and second smallest as tens, rest as ones
	return (digits[0]*10 + digits[2]) + (digits[1]*10 + digits[3])
}
```
