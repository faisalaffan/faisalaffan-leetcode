# 2231 — Largest Number After Digit Swaps By Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func LargestNumberAfterDigitSwapsByParity(num int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(log n * log(log n)), Space: O(log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2231: Largest Number After Digit Swaps by Parity
// https://leetcode.com/problems/largest-number-after-digit-swaps-by-parity/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LargestNumberAfterDigitSwapsByParity(1234)) // 3412
	fmt.Println(LargestNumberAfterDigitSwapsByParity(65875)) // 87655
}

// Time: O(log n * log(log n)), Space: O(log n)
func LargestNumberAfterDigitSwapsByParity(num int) int {
	var digits []int
	for n := num; n > 0; n /= 10 {
		digits = append(digits, n%10)
	}

	// Reverse to get original order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// Collect even and odd digits
	var evens, odds []int
	for _, d := range digits {
		if d%2 == 0 {
			evens = append(evens, d)
		} else {
			odds = append(odds, d)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(evens)))
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))

	// Reconstruct
	result := 0
	ei, oi := 0, 0
	for _, d := range digits {
		result *= 10
		if d%2 == 0 {
			result += evens[ei]
			ei++
		} else {
			result += odds[oi]
			oi++
		}
	}
	return result
}
```
