# 0004 — Median Of Two Sorted Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
// Difficulty: Hard
//
// Binary search on the smaller array to find the correct partition such that
// all left elements <= all right elements. O(log(min(m, n))) time, O(1) space.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: nums1 = [1,3], nums2 = [2] => 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	// Example 2: nums1 = [1,2], nums2 = [3,4] => 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	// Edge: all zeros
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	// Edge: one empty
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for O(log(min(m,n)))
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2 // partition position in nums1
		j := (m+n+1)/2 - i      // partition position in nums2

		// Get bounding values with sentinels for edge partitions
		maxLeftA := math.MinInt64
		if i > 0 {
			maxLeftA = nums1[i-1]
		}

		minRightA := math.MaxInt64
		if i < m {
			minRightA = nums1[i]
		}

		maxLeftB := math.MinInt64
		if j > 0 {
			maxLeftB = nums2[j-1]
		}

		minRightB := math.MaxInt64
		if j < n {
			minRightB = nums2[j]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// Found the correct partition
			if (m+n)%2 == 1 {
				return float64(max(maxLeftA, maxLeftB))
			}
			return (float64(max(maxLeftA, maxLeftB)) + float64(min(minRightA, minRightB))) / 2.0
		} else if maxLeftA > minRightB {
			// Too far right — move partition left
			right = i - 1
		} else {
			// Too far left — move partition right
			left = i + 1
		}
	}

	return 0.0 // unreachable for valid input
}
```
