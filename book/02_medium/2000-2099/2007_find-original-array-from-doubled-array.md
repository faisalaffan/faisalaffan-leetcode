# 2007 — Find Original Array From Doubled Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindOriginalArrayFromDoubledArray(changed []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2007: Find Original Array From Doubled Array
// https://leetcode.com/problems/find-original-array-from-doubled-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1, 3, 4, 2, 6, 8}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{6, 3, 0, 1}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1}))
}

// Time: O(n log n), Space: O(n)
func FindOriginalArrayFromDoubledArray(changed []int) []int {
	n := len(changed)
	if n%2 == 1 {
		return []int{}
	}

  // Sort O(n log n)
	sort.Ints(changed)
	maxVal := changed[n-1]
  // Alokasi slice
	cnt := make([]int, maxVal+1)
	for _, x := range changed {
		cnt[x]++
	}

  // Alokasi slice
	ans := make([]int, 0, n/2)
	for _, x := range changed {
		if cnt[x] == 0 {
			continue
		}
		if x*2 > maxVal || cnt[x*2] == 0 {
			return []int{}
		}
		ans = append(ans, x)
		cnt[x]--
		cnt[x*2]--
	}

	if len(ans) != n/2 {
		return []int{}
	}
	return ans
}
```
