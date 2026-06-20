# 0738 — Monotone Increasing Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func monotoneIncreasingDigits(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #738: Monotone Increasing Digits
// https://leetcode.com/problems/monotone-increasing-digits/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
}

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1

	for i < len(s) && s[i] >= s[i-1] {
		i++
	}

	if i == len(s) {
		return n
	}

	for i > 0 && s[i] < s[i-1] {
		s[i-1]--
		i--
	}

	for j := i + 1; j < len(s); j++ {
		s[j] = '9'
	}

	result, _ := strconv.Atoi(string(s))
	return result
}
```
