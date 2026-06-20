# 3954 — Sum Of Compatible Numbers In Range I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfCompatibleNumbersInRangeI(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(k)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3954: Sum of Compatible Numbers in Range I
// https://leetcode.com/problems/sum-of-compatible-numbers-in-range-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfCompatibleNumbersInRangeI(2, 3))
	fmt.Println(SumOfCompatibleNumbersInRangeI(5, 1))
}

// Time: O(k)
// Space: O(1)
func SumOfCompatibleNumbersInRangeI(n int, k int) int {
	sum := 0
	start := n - k
	if start < 1 {
		start = 1
	}
	end := n + k
	for x := start; x <= end; x++ {
		if n&x == 0 {
			sum += x
		}
	}
	return sum
}
```
