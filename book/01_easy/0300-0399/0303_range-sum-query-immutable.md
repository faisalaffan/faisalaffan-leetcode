# 0303 — Range Sum Query Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(nums []int) NumArray
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n) for init, O(1) per query  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
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
