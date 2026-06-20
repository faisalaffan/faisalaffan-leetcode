# 0201 — Bitwise And Of Numbers Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func rangeBitwiseAnd(left int, right int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(log max(left, right)), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #201: Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Difficulty: Medium
// Time: O(log max(left, right)), Space: O(1)

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
  // Two-pointer loop
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
```
