# 3723 — Maximize Sum Of Squares Of Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximizeSumOfSquaresOfDigits(num int, total int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3723: Maximize Sum of Squares of Digits
// https://leetcode.com/problems/maximize-sum-of-squares-of-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func maximizeSumOfSquaresOfDigits(num int, total int) string {
	if num*9 < total {
		return ""
	}
	nines := total / 9
	rem := total % 9
	var sb strings.Builder
	sb.WriteString(strings.Repeat("9", nines))
	if rem > 0 {
		sb.WriteByte(byte(rem) + '0')
	}
	for sb.Len() < num {
		sb.WriteByte('0')
	}
	return sb.String()
}

func main() {
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 3))
	fmt.Println(maximizeSumOfSquaresOfDigits(2, 17))
	fmt.Println(maximizeSumOfSquaresOfDigits(1, 10))
}
```
