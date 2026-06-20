# 3134 — Find The Median Of The Uniqueness Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func medianOfUniquenessArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3134: Find the Median of the Uniqueness Array
// https://leetcode.com/problems/find-the-median-of-the-uniqueness-array/
// Difficulty: Hard
//
// The uniqueness array of nums is an array of distinct-counts for all subarrays.
// Find the median of this sorted uniqueness array.
// Binary search on the answer + sliding window count of subarrays with distinct count <= mid.

import (
	"fmt"
)

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	medianPos := (total + 1) / 2

	left, right := 1, n
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right) / 2
		if int(countLE(nums, mid)) >= medianPos {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func countLE(nums []int, k int) int64 {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	distinct := 0
	var count int64 = 0
	left := 0

	for right := 0; right < n; right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			distinct++
		}
		for distinct > k {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinct--
			}
			left++
		}
		count += int64(right - left + 1)
	}
	return count
}

func main() {
	// Test case 1
	nums := []int{1, 2, 3}
	fmt.Println("Test 1:", medianOfUniquenessArray(nums))
	// Expected: 1

	// Test case 2
	nums2 := []int{3, 4, 3, 4, 5}
	fmt.Println("Test 2:", medianOfUniquenessArray(nums2))
	// Expected: ?

	// Test case 3: single element
	nums3 := []int{1}
	fmt.Println("Test 3:", medianOfUniquenessArray(nums3))
	// Expected: 1

	// Test case 4: all same
	nums4 := []int{5, 5, 5, 5}
	fmt.Println("Test 4:", medianOfUniquenessArray(nums4))
	// Expected: 1

	// Test case 5: all distinct
	nums5 := []int{1, 2, 3, 4}
	fmt.Println("Test 5:", medianOfUniquenessArray(nums5))
	// Expected: 2 (subarrays: 10 total, median pos 5th/6th → min distinct count covering >=5 subarrays)
}
```
