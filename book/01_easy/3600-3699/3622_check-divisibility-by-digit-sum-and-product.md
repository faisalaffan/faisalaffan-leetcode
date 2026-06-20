# 3622 — Check Divisibility By Digit Sum And Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckDivisibilityByDigitSumAndProduct(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3622: Check Divisibility by Digit Sum and Product
// https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(99))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(23))
	fmt.Println(CheckDivisibilityByDigitSumAndProduct(10))
}

// Time: O(log n)
// Space: O(1)
func CheckDivisibilityByDigitSumAndProduct(n int) bool {
	x := n
	sum := 0
	prod := 1
	for x > 0 {
		d := x % 10
		sum += d
		prod *= d
		x /= 10
	}
	return n%(sum+prod) == 0
}
```
