# 3452 — Sum Of Good Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfGoodNumbers(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3452: Sum of Good Numbers
// https://leetcode.com/problems/sum-of-good-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfGoodNumbers([]int{1, 3, 2, 1, 5, 4}, 2))
	fmt.Println(SumOfGoodNumbers([]int{2, 1}, 1))
}

// SumOfGoodNumbers returns sum of numbers that are greater than both nums[i-k] and nums[i+k] (or if out of bounds).
// Time: O(n). Space: O(1).
func SumOfGoodNumbers(nums []int, k int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		good := true
		if i-k >= 0 && nums[i] <= nums[i-k] {
			good = false
		}
		if i+k < n && nums[i] <= nums[i+k] {
			good = false
		}
		if good {
			sum += nums[i]
		}
	}
	return sum
}
```
