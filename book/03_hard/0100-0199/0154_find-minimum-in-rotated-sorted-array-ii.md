# 0154 — Find Minimum In Rotated Sorted Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMin(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #154: Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// Minimum is in the left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right], cannot determine, shrink
			right--
		}
	}

	return nums[left]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	result := findMin(nums)
	expected := 0

	fmt.Printf("findMin(%v) = %d\n", nums, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
