# 0633 — Sum Of Square Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func JudgeSquareSum(c int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(sqrt(c))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #633: Sum of Square Numbers
// https://leetcode.com/problems/sum-of-square-numbers/
// Difficulty: Medium
// Time: O(sqrt(c))
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(JudgeSquareSum(5))
	fmt.Println(JudgeSquareSum(3))
	fmt.Println(JudgeSquareSum(4))
	fmt.Println(JudgeSquareSum(2))
}

func JudgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
```
