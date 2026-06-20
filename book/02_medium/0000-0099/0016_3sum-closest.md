# 0016 — 3Sum Closest

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func threeSumClosest(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #16: 3Sum Closest
// https://leetcode.com/problems/3sum-closest/
// Difficulty: Medium

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
  // Two-pointer: gerakkan kiri atau kanan
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}

	return closest
}

func main() {
	// Test case 1
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2

	// Test case 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1)) // 0

	// Test case 3
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100)) // 2
}

// Time: O(n^2) | Space: O(1)
```
