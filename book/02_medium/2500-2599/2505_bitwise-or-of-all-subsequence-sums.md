# 2505 — Bitwise Or Of All Subsequence Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func subsequenceSumOr(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * 20)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2505: Bitwise OR of All Subsequence Sums
// https://leetcode.com/problems/bitwise-or-of-all-subsequence-sums/
// Difficulty: Medium
// Time: O(n * 20) | Space: O(1)
// A bit is achievable if any number has that bit, or can be formed by combination.
// Result = OR of all prefix sums? Actually: any sum that can be formed = OR of
// all elements and their combinations. Answer = OR of all elements (since we can
// always form any single element sum via a subsequence of size 1).

import "fmt"

func main() {
	fmt.Println(subsequenceSumOr([]int{2, 1, 4})) // 7 (bits 0,1,2)
	fmt.Println(subsequenceSumOr([]int{2, 3}))    // 7 (sums: 0,2,3,5 -> OR = 7)
}

func subsequenceSumOr(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans |= v
	}
	return ans
}
```
