# 3551 — Minimum Swaps To Sort By Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func digitSum(x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3551: Minimum Swaps to Sort by Digit Sum
// https://leetcode.com/problems/minimum-swaps-to-sort-by-digit-sum/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func digitSum(x int) int {
	s := 0
	if x < 0 {
		x = -x
	}
	for x > 0 {
		s += x % 10
		x /= 10
	}
	return s
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumSwapsToSortByDigitSum([]int{15, 21, 3}))
	// Test case 2
	fmt.Println("Test 2:", MinimumSwapsToSortByDigitSum([]int{1, 2, 3, 4}))
	// Test case 3
	fmt.Println("Test 3:", MinimumSwapsToSortByDigitSum([]int{10, 20, 30}))
}

func MinimumSwapsToSortByDigitSum(nums []int) int {
	n := len(nums)
	type pair struct {
		val  int
		sum  int
		idx  int
	}
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{v, digitSum(v), i}
	}
  // Custom sort dengan comparator
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].sum != arr[j].sum {
			return arr[i].sum < arr[j].sum
		}
		return arr[i].val < arr[j].val
	})
	// Count swaps needed using cycle detection
	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || arr[i].idx == i {
			continue
		}
		cycleLen := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = arr[j].idx
			cycleLen++
		}
		swaps += cycleLen - 1
	}
	return swaps
}
```
