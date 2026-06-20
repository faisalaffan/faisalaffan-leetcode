# 2544 — Alternating Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func AlternatingDigitSum(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2544: Alternating Digit Sum
// https://leetcode.com/problems/alternating-digit-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import "fmt"

func main() {
	fmt.Println(AlternatingDigitSum(521)) // 4
	fmt.Println(AlternatingDigitSum(111)) // 1
	fmt.Println(AlternatingDigitSum(886996)) // 0
}

func AlternatingDigitSum(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	// Reverse to get original order
	sum := 0
	sign := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * sign
		sign = -sign
	}
	return sum
}
```
