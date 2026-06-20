# 3536 — Maximum Product Of Two Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MaximumProductOfTwoDigits(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

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
  // Linear scan O(n)
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
