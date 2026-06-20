# 0258 — Add Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddDigits(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #258: Add Digits
// https://leetcode.com/problems/add-digits/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func AddDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	fmt.Println(AddDigits(38))
	fmt.Println(AddDigits(0))
	fmt.Println(AddDigits(9))
}
```
