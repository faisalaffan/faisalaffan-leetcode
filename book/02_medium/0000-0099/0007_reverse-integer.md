# 0007 — Reverse Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func reverse(x int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log₁₀(n))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #7: Reverse Integer
// https://leetcode.com/problems/reverse-integer/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(reverse(123)) // 321

	// Test case 2
	fmt.Println(reverse(-123)) // -321

	// Test case 3
	fmt.Println(reverse(1534236469)) // 0 (overflow)
}

// Time: O(log₁₀(n)) | Space: O(1)
```
