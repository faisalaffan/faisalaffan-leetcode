# 1780 — Check If Number Is A Sum Of Powers Of Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkPowersOfThree(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log_3 n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1780: Check if Number is a Sum of Powers of Three
// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Difficulty: Medium
// Time: O(log_3 n), Space: O(1)

import "fmt"

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	fmt.Println(checkPowersOfThree(12)) // Expected: true (3^2 + 3^1 = 9 + 3)
	fmt.Println(checkPowersOfThree(91)) // Expected: true (3^4 + 3^2 + 3^0 = 81 + 9 + 1)
	fmt.Println(checkPowersOfThree(21)) // Expected: false (21 = 2*9 + 3, cannot use 2)
}
```
