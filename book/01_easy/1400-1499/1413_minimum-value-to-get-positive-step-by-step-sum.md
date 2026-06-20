# 1413 — Minimum Value To Get Positive Step By Step Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minStartValue(nums []int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1413: Minimum Value to Get Positive Step by Step Sum
// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// Difficulty: Easy
//
// LeetCode submission: func minStartValue(nums []int) int

import "fmt"

func main() {
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{-3, 2, -3, 4, 2})) // 5
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, 2}))             // 1
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, -2, -3}))        // 5
}

// Time: O(n), Space: O(1)
func MinimumValueToGetPositiveStepByStepSum(nums []int) int {
	minSum, sum := 0, 0
	for _, v := range nums {
		sum += v
		if sum < minSum {
			minSum = sum
		}
	}
	return -minSum + 1
}
```
