# 2160 — Minimum Sum Of Four Digit Number After Splitting Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumSumOfFourDigitNumberAfterSplittingDigits(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2160: Minimum Sum of Four Digit Number After Splitting Digits
// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(2932)) // 52
	fmt.Println(MinimumSumOfFourDigitNumberAfterSplittingDigits(4009)) // 13
}

// Time: O(1), Space: O(1)
func MinimumSumOfFourDigitNumberAfterSplittingDigits(num int) int {
  // Alokasi slice integer
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = num % 10
		num /= 10
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(digits)
	// Smallest sum: smallest digit and second smallest as tens, rest as ones
	return (digits[0]*10 + digits[2]) + (digits[1]*10 + digits[3])
}
```
