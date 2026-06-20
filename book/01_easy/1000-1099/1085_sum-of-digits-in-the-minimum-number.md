# 1085 — Sum Of Digits In The Minimum Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfDigits(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1085: Sum of Digits in the Minimum Number
// https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8})) // 0 (min=1, sum=1, odd)
	fmt.Println(sumOfDigits([]int{99, 77, 33, 66, 55}))           // 1 (min=33, sum=6, even)
}

// LeetCode submission: sumOfDigits
func sumOfDigits(nums []int) int {
	x := nums[0]
	for _, v := range nums {
		if v < x {
			x = v
		}
	}
	s := 0
	for x > 0 {
		s += x % 10
		x /= 10
	}
	if s%2 == 0 {
		return 1
	}
	return 0
}
```
