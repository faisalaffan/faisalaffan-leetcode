# 0556 — Next Greater Element Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NextGreaterElementIii(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(log n) = O(number of digits)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #556: Next Greater Element III
// https://leetcode.com/problems/next-greater-element-iii/
// Difficulty: Medium
// Time: O(log n) = O(number of digits)
// Space: O(log n)

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(NextGreaterElementIii(12))
	fmt.Println(NextGreaterElementIii(21))
	fmt.Println(NextGreaterElementIii(1234))
}

func NextGreaterElementIii(n int) int {
	s := []byte(strconv.Itoa(n))

	// Find first decreasing digit from right
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	// Find smallest digit larger than s[i] from right
	j := len(s) - 1
	for s[j] <= s[i] {
		j--
	}

	s[i], s[j] = s[j], s[i]

	// Reverse suffix
	left, right := i+1, len(s)-1
  // Two-pointer loop
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	result, _ := strconv.Atoi(string(s))
	if result > math.MaxInt32 {
		return -1
	}
	return result
}
```
