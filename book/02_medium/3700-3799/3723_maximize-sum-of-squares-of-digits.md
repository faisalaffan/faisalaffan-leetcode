# 3723 — Maximize Sum Of Squares Of Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximizeSumOfSquaresOfDigits(num int, total int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3723: Maximize Sum of Squares of Digits
// https://leetcode.com/problems/maximize-sum-of-squares-of-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func maximizeSumOfSquaresOfDigits(num int, total int) string {
	if num*9 < total {
		return ""
	}
	nines := total / 9
	rem := total % 9
	var sb strings.Builder
	sb.WriteString(strings.Repeat("9", nines))
	if rem > 0 {
		sb.WriteByte(byte(rem) + '0')
	}
	for sb.Len() < num {
		sb.WriteByte('0')
	}
	return sb.String()
}

func main() {
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 3))
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 17))
	fmt.Println(maximizeSumOfSquaresOfDigits(1, 10))
}
```
