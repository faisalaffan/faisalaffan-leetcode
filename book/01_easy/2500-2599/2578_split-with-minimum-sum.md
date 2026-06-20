# 2578 — Split With Minimum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SplitWithMinimumSum(num int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2578: Split With Minimum Sum
// https://leetcode.com/problems/split-with-minimum-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitWithMinimumSum(4325)) // 59
	fmt.Println(SplitWithMinimumSum(687))  // 75
}

func SplitWithMinimumSum(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
  // Sort O(n log n)
	sort.Ints(digits)

	num1, num2 := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(digits); i++ {
		if i%2 == 0 {
			num1 = num1*10 + digits[i]
		} else {
			num2 = num2*10 + digits[i]
		}
	}
	return num1 + num2
}
```
