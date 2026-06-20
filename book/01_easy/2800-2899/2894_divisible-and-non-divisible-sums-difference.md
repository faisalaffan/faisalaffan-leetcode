# 2894 — Divisible And Non Divisible Sums Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func DivisibleAndNonDivisibleSumsDifference(n int, m int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2894: Divisible and Non-divisible Sums Difference
// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: differenceOfSums
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(10, 3)) // 19
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 6))  // 15
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 1))  // -15
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: differenceOfSums
func DivisibleAndNonDivisibleSumsDifference(n int, m int) int {
	num1 := 0 // sum of numbers not divisible by m
	num2 := 0 // sum of numbers divisible by m
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			num2 += i
		} else {
			num1 += i
		}
	}
	return num1 - num2
}
```
