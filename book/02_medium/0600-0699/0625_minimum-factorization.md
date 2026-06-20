# 0625 — Minimum Factorization

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SmallestFactorization(num int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #625: Minimum Factorization
// https://leetcode.com/problems/minimum-factorization/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SmallestFactorization(48))
	fmt.Println(SmallestFactorization(15))
	fmt.Println(SmallestFactorization(1))
}

func SmallestFactorization(num int) int {
	if num < 2 {
		return num
	}

	// Build number from right to left using digits 9..2
	result := 0
	multiplier := 1

	for i := 9; i >= 2; i-- {
		for num%i == 0 {
			result += i * multiplier
			if result > math.MaxInt32 {
				return 0
			}
			multiplier *= 10
			num /= i
		}
	}

	if num > 1 {
		return 0
	}
	return result
}
```
