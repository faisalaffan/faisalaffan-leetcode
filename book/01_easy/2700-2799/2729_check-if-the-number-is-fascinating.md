# 2729 — Check If The Number Is Fascinating

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CheckIfTheNumberIsFascinating(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2729: Check if The Number is Fascinating
// https://leetcode.com/problems/check-if-the-number-is-fascinating/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(CheckIfTheNumberIsFascinating(192))
	fmt.Println(CheckIfTheNumberIsFascinating(100))
}

func CheckIfTheNumberIsFascinating(n int) bool {
	concat := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(concat) != 9 {
		return false
	}

	digits := []byte(concat)
  // Custom sort
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	return string(digits) == "123456789"
}
```
