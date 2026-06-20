# 3536 — Maximum Product Of Two Digits

## Deskripsi

**Soal:** [3536. Maximum Product Of Two Digits](https://leetcode.com/problems/maximum-product-of-two-digits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3536: Maximum Product of Two Digits
// https://leetcode.com/problems/maximum-product-of-two-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoDigits(34))
	fmt.Println(MaximumProductOfTwoDigits(10))
	fmt.Println(MaximumProductOfTwoDigits(99))
}

// MaximumProductOfTwoDigits returns the maximum product of any two digits in n.
// Time: O(log n). Space: O(1).
func MaximumProductOfTwoDigits(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	maxProd := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			prod := digits[i] * digits[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
```
