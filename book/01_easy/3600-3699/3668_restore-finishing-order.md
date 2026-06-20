# 3668 — Restore Finishing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RestoreFinishingOrder(order []int, friends []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3668: Restore Finishing Order
// https://leetcode.com/problems/restore-finishing-order/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RestoreFinishingOrder([]int{3, 1, 2, 4, 5}, []int{2, 4}))
	fmt.Println(RestoreFinishingOrder([]int{1, 2, 3, 4, 5}, []int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func RestoreFinishingOrder(order []int, friends []int) []int {
  // HashMap: O(1) lookup
	friendSet := make(map[int]bool)
	for _, f := range friends {
		friendSet[f] = true
	}

  // Alokasi slice
	res := make([]int, 0, len(friends))
	for _, id := range order {
		if friendSet[id] {
			res = append(res, id)
		}
	}
	return res
}

func init() {
	_ = sort.Ints
}
```
