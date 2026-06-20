# 3370 — Smallest Number With All Set Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SmallestNumberWithAllSetBits(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3370: Smallest Number With All Set Bits
// https://leetcode.com/problems/smallest-number-with-all-set-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestNumberWithAllSetBits(5))
	fmt.Println(SmallestNumberWithAllSetBits(10))
	fmt.Println(SmallestNumberWithAllSetBits(3))
}

// SmallestNumberWithAllSetBits returns the smallest number >= n whose binary representation consists of all 1s.
// Time: O(log n). Space: O(1).
func SmallestNumberWithAllSetBits(n int) int {
	result := 1
	for result < n {
		result = (result << 1) | 1
	}
	return result
}
```
