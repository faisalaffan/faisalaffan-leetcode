# 2909 — Minimum Sum Of Mountain Triplets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2909: Minimum Sum of Mountain Triplets II
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
	fmt.Println(minimumSum([]int{6, 5, 4, 3, 4, 5}))
}

func minimumSum(nums []int) int {
	n := len(nums)
	const inf = 1 << 30
  // Alokasi slice integer
	right := make([]int, n+1)
	right[n] = inf
	for i := n - 1; i >= 0; i-- {
		if right[i+1] < nums[i] {
			right[i] = right[i+1]
		} else {
			right[i] = nums[i]
		}
	}
	ans, left := inf, inf
	for i, x := range nums {
		if left < x && right[i+1] < x {
			sum := left + x + right[i+1]
			if sum < ans {
				ans = sum
			}
		}
		if x < left {
			left = x
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}
```
