# 2859 — Sum Of Values At Indices With K Set Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfValuesAtIndicesWithKSetBits(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2859: Sum of Values at Indices With K Set Bits
// https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/
// Difficulty: Easy
// Time: O(n * log n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(SumOfValuesAtIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
	fmt.Println(SumOfValuesAtIndicesWithKSetBits([]int{4, 3, 2, 1}, 2))
}

func SumOfValuesAtIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, v := range nums {
		if bits.OnesCount(uint(i)) == k {
			sum += v
		}
	}
	return sum
}
```
