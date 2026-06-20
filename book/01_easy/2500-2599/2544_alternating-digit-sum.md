# 2544 — Alternating Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AlternatingDigitSum(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2544: Alternating Digit Sum
// https://leetcode.com/problems/alternating-digit-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import "fmt"

func main() {
	fmt.Println(AlternatingDigitSum(521)) // 4
	fmt.Println(AlternatingDigitSum(111)) // 1
	fmt.Println(AlternatingDigitSum(886996)) // 0
}

func AlternatingDigitSum(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	// Reverse to get original order
	sum := 0
	sign := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * sign
		sign = -sign
	}
	return sum
}
```
