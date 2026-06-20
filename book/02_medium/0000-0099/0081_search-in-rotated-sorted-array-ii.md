# 0081 — Search In Rotated Sorted Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func search(nums []int, target int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(log n) average, O(n) worst  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #81: Search in Rotated Sorted Array II
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// Handle duplicates: shrink window
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// Left half is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	// Test case 1
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true

	// Test case 2
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false

	// Test case 3
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0)) // true
}

// Time: O(log n) average, O(n) worst | Space: O(1)
```
