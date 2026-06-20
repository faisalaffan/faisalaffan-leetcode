# 3645 — Maximum Total From Optimal Activation Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumTotalFromOptimalActivationOrder(value []int, limit []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3645: Maximum Total from Optimal Activation Order
// https://leetcode.com/problems/maximum-total-from-optimal-activation-order/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumTotalFromOptimalActivationOrder(value []int, limit []int) int64 {
	n := len(value)
  // Membuat matriks/slice 2D untuk DP
	groups := make([][]int, n+1)
	for i := 0; i < n; i++ {
		l := limit[i]
		groups[l] = append(groups[l], value[i])
	}

	var total int64 = 0
	for l := 1; l <= n; l++ {
		if len(groups[l]) == 0 {
			continue
		}
  // Custom sort dengan comparator
		sort.Slice(groups[l], func(i, j int) bool {
			return groups[l][i] > groups[l][j]
		})
		cap := l
		if len(groups[l]) < cap {
			cap = len(groups[l])
		}
		for i := 0; i < cap; i++ {
			total += int64(groups[l][i])
		}
	}
	return total
}

func main() {
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{3, 5, 2, 4}, []int{2, 1, 3, 2}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{10, 20}, []int{1, 1}))
	fmt.Println(maximumTotalFromOptimalActivationOrder([]int{1, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3}))
}
```
