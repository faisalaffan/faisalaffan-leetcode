# 0259 — 3Sum Smaller

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func threeSumSmaller(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #259: 3Sum Smaller
// https://leetcode.com/problems/3sum-smaller/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(1)

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	count := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
  // Two-pointer: gerakkan kiri atau kanan
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{1, 1, -2}, 1))
	fmt.Println(threeSumSmaller([]int{0, 0, 0}, 0))
}
```
