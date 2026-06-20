# 1438 — Longest Continuous Subarray With Absolute Diff Less Than Or Equal To Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubarray(nums []int, limit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n) where n = len(nums)  
**Kompleksitas Ruang:** O(n) for deques

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1438: Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4)) // 2

	// Test case 2
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5)) // 4

	// Test case 3
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0)) // 3

	// Test case 4
	fmt.Println(longestSubarray([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4)) // 5
}

// Time: O(n) where n = len(nums)
// Space: O(n) for deques
func longestSubarray(nums []int, limit int) int {
	// Monotonic deques for tracking min and max in current window
  // Alokasi slice integer
	minDeque := make([]int, 0) // increasing
  // Alokasi slice integer
	maxDeque := make([]int, 0) // decreasing

	left := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		// Maintain min deque (increasing)
		for len(minDeque) > 0 && minDeque[len(minDeque)-1] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, nums[right])

		// Maintain max deque (decreasing)
		for len(maxDeque) > 0 && maxDeque[len(maxDeque)-1] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, nums[right])

		// Shrink window if diff > limit
		for maxDeque[0]-minDeque[0] > limit {
			if nums[left] == minDeque[0] {
				minDeque = minDeque[1:]
			}
			if nums[left] == maxDeque[0] {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
```
