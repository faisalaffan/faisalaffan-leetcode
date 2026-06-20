# 2177 — Find Three Consecutive Integers That Sum To A Given Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfThree(num int64) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2177: Find Three Consecutive Integers That Sum to a Given Number
// https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	x := num / 3
	return []int64{x - 1, x, x + 1}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfThree(33))
	// Expected: [10, 11, 12]

	// Test case 2
	fmt.Println("Test 2:", sumOfThree(4))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", sumOfThree(0))
	// Expected: [-1, 0, 1]
}
```
