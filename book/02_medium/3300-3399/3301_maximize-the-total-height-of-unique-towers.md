# 3301 — Maximize The Total Height Of Unique Towers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTotalSum(maximumHeight []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3301: Maximize the Total Height of Unique Towers
// https://leetcode.com/problems/maximize-the-total-height-of-unique-towers/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTotalSum([]int{2, 3, 4, 3}))   // 10
	fmt.Println(maximumTotalSum([]int{2, 2, 1}))      // -1
	fmt.Println(maximumTotalSum([]int{5, 4, 3, 2, 1})) // 15
}

func maximumTotalSum(maximumHeight []int) int64 {
  // Custom sort dengan comparator
	sort.Slice(maximumHeight, func(i, j int) bool {
		return maximumHeight[i] > maximumHeight[j]
	})

	var total int64
	prev := maximumHeight[0]
	total += int64(prev)

	for i := 1; i < len(maximumHeight); i++ {
		if prev <= 1 {
			return -1
		}
		h := maximumHeight[i]
		if h >= prev {
			h = prev - 1
		}
		total += int64(h)
		prev = h
	}
	return total
}
```
