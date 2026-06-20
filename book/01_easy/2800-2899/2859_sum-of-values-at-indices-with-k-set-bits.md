# 2859 — Sum Of Values At Indices With K Set Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SumOfValuesAtIndicesWithKSetBits(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * log n)  |  **Ruang:** O(1)


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
