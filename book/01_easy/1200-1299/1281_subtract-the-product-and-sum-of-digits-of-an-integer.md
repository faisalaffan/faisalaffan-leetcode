# 1281 — Subtract The Product And Sum Of Digits Of An Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func subtractProductAndSum(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1281: Subtract the Product and Sum of Digits of an Integer
// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Difficulty: Easy
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))  // 15
	fmt.Println(subtractProductAndSum(4421)) // 21
}

// LeetCode submission: subtractProductAndSum
func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		product *= d
		sum += d
	}
	return product - sum
}
```
