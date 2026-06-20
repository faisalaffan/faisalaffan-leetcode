# 0303 — Range Sum Query Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(nums []int) NumArray`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n) for init, O(1) per query  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #303: Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/
// Difficulty: Easy

import "fmt"

type NumArray struct {
	prefix []int
}

// Time: O(n) for init, O(1) per query | Space: O(n)
func Constructor(nums []int) NumArray {
  // Alokasi slice
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	return NumArray{prefix: prefix}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
```
