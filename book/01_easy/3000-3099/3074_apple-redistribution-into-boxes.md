# 3074 — Apple Redistribution Into Boxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AppleRedistributionIntoBoxes(apples []int, capacity []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3074: Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumBoxes
	fmt.Println(AppleRedistributionIntoBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2})) // 2
	fmt.Println(AppleRedistributionIntoBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))    // 4
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumBoxes
func AppleRedistributionIntoBoxes(apples []int, capacity []int) int {
	totalApples := 0
	for _, a := range apples {
		totalApples += a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	boxes := 0
	for _, c := range capacity {
		boxes++
		totalApples -= c
		if totalApples <= 0 {
			return boxes
		}
	}
	return boxes
}
```
