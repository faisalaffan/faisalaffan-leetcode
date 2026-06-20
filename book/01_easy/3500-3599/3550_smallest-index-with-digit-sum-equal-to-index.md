# 3550 — Smallest Index With Digit Sum Equal To Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumDigits(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * log max). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3550: Smallest Index With Digit Sum Equal to Index
// https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SmallestIndexWithDigitSumEqualToIndex([]int{10, 11, 12, 13, 14}))
}

// digitSum returns sum of digits.
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// SmallestIndexWithDigitSumEqualToIndex returns the smallest index i where digit sum of nums[i] equals i, else -1.
// Time: O(n * log max). Space: O(1).
func SmallestIndexWithDigitSumEqualToIndex(nums []int) int {
	for i, v := range nums {
		if sumDigits(v) == i {
			return i
		}
	}
	return -1
}
```
