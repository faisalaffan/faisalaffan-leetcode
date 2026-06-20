# 0829 — Consecutive Numbers Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func consecutiveNumbersSum(n int) int
```

> **💡 Hint:** For each possible length m, check if N = m*k + m*(m-1)/2.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #829: Consecutive Numbers Sum
// https://leetcode.com/problems/consecutive-numbers-sum/
// Difficulty: Hard
// Approach: For each possible length m, check if N = m*k + m*(m-1)/2.
// Which simplifies to: (N - m*(m-1)/2) % m == 0 and > 0.

import "fmt"

func consecutiveNumbersSum(n int) int {
	count := 0
	for m := 1; m*(m-1)/2 < n; m++ {
		rem := n - m*(m-1)/2
		if rem%m == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(consecutiveNumbersSum(5))  // Expected: 2
	fmt.Println(consecutiveNumbersSum(9))  // Expected: 3
	fmt.Println(consecutiveNumbersSum(15)) // Additional test: 15 = 15, 7+8, 4+5+6, 1+2+3+4+5 => 4
}
```
