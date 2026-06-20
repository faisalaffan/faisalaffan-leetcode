# 3353 — Minimum Total Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumTotalOperations(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3353: Minimum Total Operations
// https://leetcode.com/problems/minimum-total-operations/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MinimumTotalOperations([]int{1, 2, 3, 4}))
	fmt.Println(MinimumTotalOperations([]int{1, 1, 1}))
}

// MinimumTotalOperations returns the minimum number of operations to make all elements zero.
// Each operation picks a subarray and subtracts the minimum value in it from all its elements.
// Time: O(n). Space: O(1).
func MinimumTotalOperations(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
	ops := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ops += nums[i] - nums[i-1]
		}
	}
	return ops
}
```
