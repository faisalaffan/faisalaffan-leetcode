# 3754 — Concatenate Non Zero Digits And Multiply By Sum I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func ConcatenateNonZeroDigitsAndMultiplyBySumI(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3754: Concatenate Non-Zero Digits and Multiply by Sum I
// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenateNonZeroDigitsAndMultiplyBySumI(10203004))
	fmt.Println(ConcatenateNonZeroDigitsAndMultiplyBySumI(1000))
}

// Time: O(log n)
// Space: O(1)
func ConcatenateNonZeroDigitsAndMultiplyBySumI(n int) int {
	concat := 0
	digitSum := 0
	multiplier := 1

	for n > 0 {
		d := n % 10
		if d != 0 {
			concat = d*multiplier + concat
			multiplier *= 10
			digitSum += d
		}
		n /= 10
	}

	return concat * digitSum
}
```
