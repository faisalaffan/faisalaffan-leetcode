# 0493 — Reverse Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func reversePairs(nums []int) int
```

> **💡 Hint:** Merge sort counting. During merge, count pairs (i, j) where

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #493: Reverse Pairs
// https://leetcode.com/problems/reverse-pairs/
// Difficulty: Hard
// Approach: Merge sort counting. During merge, count pairs (i, j) where
// i < j and nums[i] > 2 * nums[j]. This is similar to counting inversions
// but with the 2x multiplier.

import (
	"fmt"
)

func main() {
	fmt.Println("493 - Reverse Pairs")

	// Test cases
	fmt.Printf("reversePairs([1,3,2,3,1]) = %d (expected: 2)\n",
		reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Printf("reversePairs([2,4,3,5,1]) = %d (expected: 3)\n",
		reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Printf("reversePairs([]) = %d (expected: 0)\n",
		reversePairs([]int{}))
	fmt.Printf("reversePairs([1]) = %d (expected: 0)\n",
		reversePairs([]int{1}))
	fmt.Printf("reversePairs([1,1,1,1]) = %d (expected: 0)\n",
		reversePairs([]int{1, 1, 1, 1}))
	fmt.Printf("reversePairs([5,4,3,2,1]) = %d (expected: 4)\n",
		reversePairs([]int{5, 4, 3, 2, 1}))
	fmt.Printf("reversePairs([2147483647, 2147483647, -2147483648, -2147483648]) = %d\n",
		reversePairs([]int{2147483647, 2147483647, -2147483648, -2147483648}))
	fmt.Printf("reversePairs([1,2,3,4,5]) = %d (expected: 0)\n",
		reversePairs([]int{1, 2, 3, 4, 5}))
}

func reversePairs(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
  // Alokasi slice integer
	temp := make([]int, len(nums))
	count := mergeSort(nums, temp, 0, len(nums)-1)
	return count
}

func mergeSort(nums []int, temp []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	count := mergeSort(nums, temp, left, mid)
	count += mergeSort(nums, temp, mid+1, right)

	// Count reverse pairs across two halves
	count += countPairs(nums, left, mid, right)

	// Merge the two sorted halves
	merge(nums, temp, left, mid, right)

	return count
}

// countPairs counts how many reverse pairs exist between [left..mid] and [mid+1..right]
// Both halves are sorted.
func countPairs(nums []int, left, mid, right int) int {
	count := 0
	j := mid + 1

	for i := left; i <= mid; i++ {
		// For each nums[i], find first nums[j] where nums[i] <= 2*nums[j]
		// Use int64 to avoid overflow
		val := int64(nums[i])
		for j <= right && val > 2*int64(nums[j]) {
			j++
		}
		count += j - (mid + 1)
	}

	return count
}

func merge(nums []int, temp []int, left, mid, right int) {
	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}

	for idx := left; idx <= right; idx++ {
		nums[idx] = temp[idx]
	}
}
```
