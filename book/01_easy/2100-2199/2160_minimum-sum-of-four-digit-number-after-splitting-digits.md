# 2160 — Minimum Sum Of Four Digit Number After Splitting Digits

## Deskripsi

**Soal:** [2160. Minimum Sum Of Four Digit Number After Splitting Digits](https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = num % 10
		num /= 10
	}
	sort.Ints(digits)
	// Smallest sum: smallest digit and second smallest as tens, rest as ones
	return (digits[0]*10 + digits[2]) + (digits[1]*10 + digits[3])
}
```
