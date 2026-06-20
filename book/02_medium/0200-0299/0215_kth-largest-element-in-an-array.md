# 0215 — Kth Largest Element In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findKthLargest(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n) average with quickselect, O(n log n) worst case, Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #215: Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
// Difficulty: Medium
// Time: O(n) average with quickselect, O(n log n) worst case, Space: O(log n)

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kSmallest int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := partition(nums, left, right)

	if kSmallest == pivotIdx {
		return nums[kSmallest]
	} else if kSmallest < pivotIdx {
		return quickSelect(nums, left, pivotIdx-1, kSmallest)
	}
	return quickSelect(nums, pivotIdx+1, right, kSmallest)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	storeIdx := left

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}

	nums[storeIdx], nums[right] = nums[right], nums[storeIdx]
	return storeIdx
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{1}, 1))
}
```
