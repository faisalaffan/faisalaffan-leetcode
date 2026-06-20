# 1837 — Sum Of Digits In Base K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumBase(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log_k(n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1837: Sum of Digits in Base K
// https://leetcode.com/problems/sum-of-digits-in-base-k/
// Difficulty: Easy

import "fmt"

// Time: O(log_k(n)), Space: O(1)
func SumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}

func main() {
	fmt.Println(SumBase(34, 6))
	fmt.Println(SumBase(10, 10))
	fmt.Println(SumBase(42, 2))
}
```
