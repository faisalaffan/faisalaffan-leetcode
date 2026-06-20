# 0633 — Sum Of Square Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func JudgeSquareSum(c int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(sqrt(c))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #633: Sum of Square Numbers
// https://leetcode.com/problems/sum-of-square-numbers/
// Difficulty: Medium
// Time: O(sqrt(c))
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(JudgeSquareSum(5))
	fmt.Println(JudgeSquareSum(3))
	fmt.Println(JudgeSquareSum(4))
	fmt.Println(JudgeSquareSum(2))
}

func JudgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

  // Binary search loop
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
```
