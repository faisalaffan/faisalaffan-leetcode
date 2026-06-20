# 2180 — Count Integers With Even Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountIntegersWithEvenDigitSum(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2180: Count Integers With Even Digit Sum
// https://leetcode.com/problems/count-integers-with-even-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIntegersWithEvenDigitSum(4))  // 2
	fmt.Println(CountIntegersWithEvenDigitSum(30))  // 14
}

// Time: O(n log n), Space: O(1)
func CountIntegersWithEvenDigitSum(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if digitSumEven(i) {
			count++
		}
	}
	return count
}

func digitSumEven(n int) bool {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum%2 == 0
}
```
