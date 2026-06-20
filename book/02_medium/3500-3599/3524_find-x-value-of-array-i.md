# 3524 — Find X Value Of Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindXValueOfArrayI(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3524: Find X Value of Array I
// https://leetcode.com/problems/find-x-value-of-array-i/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindXValueOfArrayI([]int{1, 2, 3, 4}, 4))
	// Test case 2
	fmt.Println("Test 2:", FindXValueOfArrayI([]int{1, 1, 2, 3}, 3))
	// Test case 3
	fmt.Println("Test 3:", FindXValueOfArrayI([]int{5, 1, 3}, 2))
}

func FindXValueOfArrayI(nums []int, target int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	// Find x such that sum of (nums[i] > x ? x : nums[i]) equals target
	// Using prefix sums and binary search
  // Alokasi slice integer
	prefix := make([]int, len(nums)+1)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Binary search on x
	left, right := 0, nums[len(nums)-1]
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		// Find first index > mid
		idx := sort.Search(len(nums), func(i int) bool { return nums[i] > mid })
		sum := prefix[idx] + mid*(len(nums)-idx)
		if sum >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```
