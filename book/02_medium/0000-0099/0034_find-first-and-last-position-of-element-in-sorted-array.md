# 0034 — Find First And Last Position Of Element In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func searchRange(nums []int, target int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #34: Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Difficulty: Medium

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return result
	}

	// Find first position
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < len(nums) && nums[left] == target {
		result[0] = left
	} else {
		return result
	}

	// Find last position
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	result[1] = right

	return result
}

func main() {
	// Test case 1
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]

	// Test case 2
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1, -1]

	// Test case 3
	fmt.Println(searchRange([]int{}, 0)) // [-1, -1]
}

// Time: O(log n) | Space: O(1)
```
