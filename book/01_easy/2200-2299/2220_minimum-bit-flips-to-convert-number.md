# 2220 — Minimum Bit Flips To Convert Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumBitFlipsToConvertNumber(start int, goal int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2220: Minimum Bit Flips to Convert Number
// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumBitFlipsToConvertNumber(10, 7))  // 3
	fmt.Println(MinimumBitFlipsToConvertNumber(3, 4))   // 3
}

// Time: O(1), Space: O(1)
func MinimumBitFlipsToConvertNumber(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor > 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}
```
