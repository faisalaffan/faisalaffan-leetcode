# 2578 — Split With Minimum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SplitWithMinimumSum(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2578: Split With Minimum Sum
// https://leetcode.com/problems/split-with-minimum-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitWithMinimumSum(4325)) // 59
	fmt.Println(SplitWithMinimumSum(687))  // 75
}

func SplitWithMinimumSum(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(digits)

	num1, num2 := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(digits); i++ {
		if i%2 == 0 {
			num1 = num1*10 + digits[i]
		} else {
			num2 = num2*10 + digits[i]
		}
	}
	return num1 + num2
}
```
