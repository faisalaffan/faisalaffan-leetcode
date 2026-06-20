# 1833 — Maximum Ice Cream Bars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxIceCream(costs []int, coins int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1833: Maximum Ice Cream Bars
// https://leetcode.com/problems/maximum-ice-cream-bars/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
  // Sort O(n log n)
	sort.Ints(costs)
	count := 0
	for _, c := range costs {
		if coins >= c {
			coins -= c
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7)) // Expected: 4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) // Expected: 0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) // Expected: 6
}
```
