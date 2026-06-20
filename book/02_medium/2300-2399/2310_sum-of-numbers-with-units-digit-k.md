# 2310 — Sum Of Numbers With Units Digit K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumNumbers(num int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2310: Sum of Numbers With Units Digit K
// https://leetcode.com/problems/sum-of-numbers-with-units-digit-k/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if (i*k)%10 == num%10 && i*k <= num {
			return i
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println(minimumNumbers(58, 9))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumNumbers(37, 2))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumNumbers(0, 7))
	// Expected: 0
}
```
